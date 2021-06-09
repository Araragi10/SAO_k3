package wotoTranslate

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoTranslate/wotoLang"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TrHandler(message *tg.Message, args *pTools.EventArgs) {
	send_pv := args.GetAsBool(PV_FLAG, PRIVATE_FLAG)
	is_reply := message.ReplyToMessage != nil
	gnu := args.GetAsBool(GNU_FLAG)
	toLang := wv.EMPTY
	frLang := wv.EMPTY

	if args.IsEmptyOrRaw() && !is_reply {
		return
	}

	if is_reply && args.IsEmptyOrRaw() {
		toLang = wotoLang.L_en
	} else {
		toLang = args.GetAsStringTOrRaw(TO_FLAG)
	}

	if ws.IsEmpty(&toLang) {
		f1 := args.GetIndexFlag(wv.BaseIndex)
		if f1 != nil {
			namae := f1.GetName()
			if wotoLang.IsLang(namae) {
				toLang = namae
			} else {
				toLang = wotoLang.L_en
			}
		}
	}

	// please don't use GetAsStringTOrRaw here.
	// we don't know if what user entered is from or not.
	frLang = args.GetAsStringT(FR_FLAG, FROM_FLAG)

	directly := args.HasFlag(FR_FLAG, FROM_FLAG)
	frSet := func() {
		if ws.IsEmpty(&frLang) {
			frLang = wv.AutoStr
		}

		if !directly {
			directly = true
		}
	}

	var full string
	if is_reply {
		if !ws.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			if !ws.IsEmpty(&message.ReplyToMessage.Caption) {
				full = message.ReplyToMessage.Caption
			} else {
				// we should check if the replied message has text or not,
				// and since it doesn't have any text value, then we should not
				// do any operations on it, instead we have to get the
				// text from current message (and of course reply the
				// translation to the replied text)
				full = args.GetAsStringTOrRaw(TXT_FLAG, TEXT_FLAG)
			}
		}
	} else {
		// do not convert the flags to the morse code.
		full = args.GetAsStringTOrRaw(TXT_FLAG, TEXT_FLAG)
	}

	//log.Println(full)
	if ws.IsEmpty(&full) {
		return
	}

	var l1 *Lang
	var cl []LangDetect

	if !directly {
		l1 = DetectLanguage(full)
		if l1 == nil {
			frSet()
		} else {
			cl = l1.Data.Detections
			if len(cl) == wv.BaseIndex {
				frSet()
			}
		}
	} else {
		frSet()
	}

	var str string

	if gnu {
		if !directly {
			if l1 == nil {
				return
			}

			if l1.Data == nil {
				return
			}

			if len(l1.Data.Detections) == wv.BaseIndex {
				return
			}

			str = TrGnuTxt(l1.Data.Detections[0].TheLang, toLang, full)
		} else {
			str = TrGnuTxt(wv.AutoStr, toLang, full)
		}
	} else {
		var trl *WotoTr
		if !directly {
			trl = Translate(l1, toLang, full)
		} else {
			trl = TranslateD(frLang, toLang, full)
		}
		if trl == nil {
			log.Println("trl is nil.")
			return
		}

		if trl.HasWrongNess {
			str += "Did you mean \"" + trl.CorrectedValue + "\"?\n"
		}
		str += trl.TranslatedText
	}

	sendTr(message, &str, is_reply, send_pv)
}

// please before using this function, ensure that the `morse` value is
// converted to markdown (using wotoMarkdown) functions.
func sendTr(message *tg.Message, morse *string, reply, pv bool) {
	// always check before sending:
	// Bad Request: message text is empty
	if ws.IsEmpty(morse) {
		return
	}

	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	var msg tg.MessageConfig
	if pv {
		//settings.SendSudo(strconv.Itoa(int(message.From.ID)))
		msg = tg.NewMessage(message.From.ID, (*morse))
	} else {
		//settings.SendSudo(strconv.Itoa(int(message.Chat.ID)))
		msg = tg.NewMessage(message.Chat.ID, (*morse))
	}

	// !pv => for fixing: Bad Request: replied message not found
	if reply && !pv {
		r := message.ReplyToMessage
		if r != nil {
			msg.ReplyToMessageID = r.MessageID
		} else {
			msg.ReplyToMessageID = message.MessageID
		}
	} else {
		// for fixing: Bad Request: replied message not found
		if !pv {
			msg.ReplyToMessageID = message.MessageID
		}
	}

	//msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}
