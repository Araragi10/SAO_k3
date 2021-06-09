package wotoGP

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Araragi10/SAO_k3/wotoPacks/appSettings"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/pTools"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoMD"
	ws "github.com/Araragi10/SAO_k3/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoValues/tgMessages"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gotd/td/tg"
)

func PingHandler(message *tgbot.Message, args *pTools.EventArgs) {
	pv := args.GetAsBool(PV_FLAG, PRIVATE_FLAG)
	reply := message.ReplyToMessage != nil
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	start := time.Now()
	str := wotoMD.GetItalic(pingingMessage).ToString()

	var msg tgbot.MessageConfig
	if pv {
		//settings.SendSudo(strconv.Itoa(int(message.From.ID)))
		msg = tgbot.NewMessage(message.From.ID, str)
	} else {
		//settings.SendSudo(strconv.Itoa(int(message.Chat.ID)))
		msg = tgbot.NewMessage(message.Chat.ID, str)
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
	msg.ParseMode = tgbot.ModeMarkdownV2
	res, err := api.Send(msg)
	if err != nil {

		log.Println(err)
		tgbotErr := tgMessages.GetTgError(err)
		if tgbotErr != nil {
			tgbotErr.SendRandomErrorMessage(message)
		}
		return
	}

	str = fmt.Sprintf(pingMessage, time.Since(start))
	str = wotoMD.GetItalic(str).ToString()
	eM := tgbot.NewEditMessageText(msg.ChatID, res.MessageID, str)
	eM.ParseMode = tgbot.ModeMarkdownV2
	res, err = api.Send(eM)
	if err != nil {

		log.Println(err)
		tgbotErr := tgMessages.GetTgError(err)
		if tgbotErr != nil {
			tgbotErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func BanHandler(message *tgbot.Message, args *pTools.EventArgs) {
	if !message.Chat.IsSuperGroup() {
		log.Println("it's not a super group")
		return
	}

	settings := appSettings.GetExisting()
	if settings == nil {
		log.Println("settings is nil")
		return
	}

	api := settings.GetAPI()
	if api == nil {
		log.Println("api cannot be nil.")
		return
	}

	chat := message.Chat.ChatConfig()

	sender := GetMember(api, &chat, message.From)
	if sender == nil {
		log.Println("sender of message cannot be nil.")
		return
	}

	if !sender.IsAdministrator() && !sender.IsCreator() {
		delC := tgbot.NewDeleteMessage(chat.ChatID, message.MessageID)
		api.Request(delC)
		return
	}

	reply := message.ReplyToMessage != nil

	if !reply && args.IsEmpty() {
		return
	}

	var id int64
	sendMe := func(txt string, md bool) {
		msg := tgbot.NewMessage(message.Chat.ID, txt)
		msg.ReplyToMessageID = message.MessageID
		if md {
			msg.ParseMode = tgbot.ModeMarkdownV2
		}
		_, errM := api.Send(msg)
		if errM != nil {
			log.Println("errM: " + errM.Error())
		}
	}

	if reply {
		target := message.ReplyToMessage.From
		if target.ID == message.From.ID {
			sendMe("are you this much depressed that wanna ban yourself?", false)
			// yeah, trying to ban yourself?
			return
		} else if target.ID == api.Self.ID {
			sendMe("don't ban me baka", false)
			return
		} else {
			id = message.ReplyToMessage.From.ID
		}
	} else {
		uname := args.GetAsStringOrRaw(USER_FLAG, USR_FLAG)
		// parse it to int64, if we can't, then it means it should
		// be a username.
		iS, err := strconv.ParseInt(uname, wv.BaseTen, wv.Base64Bit)
		if err != nil {
			// since there is an error here, it means
			// it should be an username, tho we have to check and see
			// if this is a valid username or not.
			uname = strings.TrimSpace(uname)
			uname = strings.Trim(uname, wv.AtSign)
			uname = strings.TrimSpace(uname)
			if !IsValidUsername(uname) {
				log.Println("username is actually invalid: " + uname)
				return
			}
			g, ctx := settings.GetGClient()
			if g == nil || ctx == nil {
				log.Println("g cannot be nil.")
				return
			}

			r, err := g.ContactsResolveUsername(*ctx, uname)
			if err != nil {
				sendMe(err.Error(), false)
				return
			}

			// check if this user actually exists or not.
			if len(r.Users) != wv.BaseIndex {
				// it means this is really a user.
				id = int64(r.Users[wv.BaseIndex].GetID())
			} else if len(r.Chats) != wv.BaseIndex {
				// it means this id
				ch, ok := r.Chats[wv.BaseIndex].(*tg.Channel)
				if !ok {
					log.Println("couldn't parse " +
						reflect.TypeOf(r.Chats[wv.BaseIndex]).Name() +
						"to channel.")
					return
				}

				isChannel := !ch.Megagroup && !ch.Gigagroup

				var strNm string
				if isChannel {
					strNm = "channel"
				} else {
					strNm = "group"
				}
				fStr := "But " + ch.Title +
					" is a " + strNm + " and cannot be banned!"
				sendMe(fStr, false)
				return
			} else {
				log.Println("unsuppourted format:" +
					reflect.TypeOf(r).String())
				// do nothing and return
				return
			}
		} else {
			id = iS
		}

	}

	if id == wv.BaseIndex {
		log.Println("id is zero, thus we return this.")
		return
	}

	//me := GetMeInGp(api, &chat)
	//if me == nil {
	//	return
	//}

	//if !me.IsAdministrator() {
	//	return
	//} else {
	//	if !me.CanRestrictMembers {
	//		return
	//	}
	//}

	cfg := tgbot.KickChatMemberConfig{
		ChatMemberConfig: tgbot.ChatMemberConfig{
			ChatID:             chat.ChatID,
			SuperGroupUsername: chat.SuperGroupUsername,
			UserID:             id,
		},
		UntilDate:      wv.BaseIndex,
		RevokeMessages: false,
	}
	_, errS := api.Request(cfg)
	if errS != nil {
		log.Println(errS)
		return
	}

	sendMe("banned", false)
}

func InfoHandler(message *tgbot.Message, args *pTools.EventArgs) {
	// reply := message.ReplyToMessage != nil
	uname := args.GetAsStringOrRaw(USER_FLAG, USR_FLAG)
	var uId int64
	if ws.IsEmpty(&uname) {
		if message.ReplyToMessage == nil {
			return
		}

		uname = message.ReplyToMessage.From.UserName
		if ws.IsEmpty(&uname) {
			uId = message.ReplyToMessage.From.ID
			if uId == wv.BaseIndex {
				return
			}

		}
	}

	s := appSettings.GetExisting()

	api := s.GetAPI()
	if api == nil {
		return
	}

	uname = strings.TrimSpace(uname)
	for strings.HasPrefix(uname, wv.AtSign) {
		uname = strings.TrimPrefix(uname, wv.AtSign)
	}
	if strings.Contains(uname, wv.SPACE_VALUE) {
		sendErr(message, usernameWrong)
	}

	//u, err := GetUserInfo(uname)
	//if err != nil {
	//	s.SendSudo(err.Error())
	//	return
	//}

	//b, err := json.MarshalIndent(u, "", "  ")
	//if err != nil {
	//	s.SendSudo(err.Error())
	//	return
	//}

	t := GetTotalInfo(uname, true)
	if ws.IsEmpty(&t) {
		log.Println("t was empty.")
		return
	}

	pv := args.HasFlag(PRIVATE_FLAG, PV_FLAG)
	reply := message.ReplyToMessage != nil
	del := (pv || reply) &&
		args.HasFlag(DEL_FLAG, DELETE_FLAG)

	var cId int64
	if pv {
		cId = message.From.ID
	} else {
		cId = message.Chat.ID
	}

	if del {
		go func() {
			req := tgbot.NewDeleteMessage(message.Chat.ID, message.MessageID)
			api.Send(req)
		}()
	}

	msg := tgbot.NewMessage(cId, t)
	if reply {
		msg.ReplyToMessageID = message.ReplyToMessage.MessageID
	} else {
		msg.ReplyToMessageID = message.MessageID
	}

	msg.ParseMode = tgbot.ModeMarkdownV2

	if _, err := api.Send(msg); err != nil {
		tgbotErr := tgMessages.GetTgError(err)
		if tgbotErr != nil {
			tgbotErr.SendRandomErrorMessage(message)
		}
	}
}

func sendErr(message *tgbot.Message, code gpError) {

}
