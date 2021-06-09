// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

import (
	"log"

	"github.com/Araragi10/SAO_k3/wotoPacks/appSettings"
	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/pTools"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoSecurity/wotoStrings"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoValues/tgMessages"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TickleHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomTickle()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func SlapHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomSlap()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func PokeHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomPoke()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func NekoHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomNeko()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func MeowHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomMeow()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func LizardHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomLizard()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func KissHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomKiss()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func HugHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomHug()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func FoxHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomFoxGirl()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func FeedHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomFeed()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func CuddleHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomCuddle()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func KemonomimiHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomKemonomimi()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func HoloHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomHolo()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func SmugHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomSmug()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func BakaHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomBaka()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func WoofHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomWoof()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func GooseHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomGoose()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func GecgHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomGecg()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func AvatarHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomAvatar()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func WaifuHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomWaifu()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func WhyHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomWhy()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func NameHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomName()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func CatHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomCat()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func FactHandler(message *tg.Message, args *pTools.EventArgs) {
	base, err := GetRandomFact()
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

func OwoHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	base, err := GetOwo(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
}

//---------------------------------------------------------

func TGifDHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetRandomTGifD(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func Mp4Handler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomMp4(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func MediumgifHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomMediumgif(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func NanogifHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomNanogif(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func TinywebmHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomTinywebm(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func GifHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomGif(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func Tinymp4Handler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomTinymp4(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func NanowebmHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomNanowebm(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func Nanomp4Handler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomNanomp4(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func Loopedmp4Handler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomLoopedmp4(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func WebmHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomWebm(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

func TinygifHandler(message *tg.Message, args *pTools.EventArgs) {
	str := getCorrectText(message, args)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	sended, ok := checkTenor(message)
	if !ok {
		return
	}

	base, err := GetTenorRandomTinygif(str)
	if err != nil {
		log.Println(err)
		return
	}

	sendNekoBase(message, base, args)
	endTenor(sended)
}

//---------------------------------------------------------

func sendNekoBase(message *tg.Message, neko interfaces.NekoFactor, args *pTools.EventArgs) {
	if message == nil || neko == nil {
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

	pv := args.GetAsBool(PRIVATE_FLAG, PV_FLAG)
	hasMsg := args.HasFlag(MSG_FLAG, MESSAGE_FLAG)
	reply := message.ReplyToMessage != nil
	del := (reply || pv) &&
		args.GetAsBool(DEL_FLAG, DELETE_FLAG)
	text := args.GetAsStringOrRaw(MSG_FLAG, MESSAGE_FLAG)

	var sendIt tg.Chattable

	if neko.IsPhoto() {
		var photo tg.PhotoConfig
		if pv {
			photo = tg.NewPhoto(message.From.ID, tg.FileURL(neko.GetUrl()))
			if message.From.ID == message.Chat.ID {
				photo.ReplyToMessageID = message.MessageID
			}
		} else {
			photo = tg.NewPhoto(message.Chat.ID, tg.FileURL(neko.GetUrl()))
			if reply && message.ReplyToMessage != nil {
				photo.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				photo.ReplyToMessageID = message.MessageID
			}
		}
		if hasMsg {
			photo.Caption = text
		}
		sendIt = photo
	} else if neko.IsGif() {
		var gif tg.DocumentConfig
		if pv {
			gif = tg.NewDocument(message.From.ID, tg.FileURL(neko.GetUrl()))
			if message.From.ID == message.Chat.ID {
				gif.ReplyToMessageID = message.MessageID
			}
		} else {
			gif = tg.NewDocument(message.Chat.ID, tg.FileURL(neko.GetUrl()))
			if reply && message.ReplyToMessage != nil {
				gif.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				gif.ReplyToMessageID = message.MessageID
			}
		}

		if hasMsg {
			gif.Caption = text
		}
		sendIt = gif
	} else if neko.IsText() {
		var msg tg.MessageConfig
		if pv {
			msg = tg.NewMessage(message.From.ID, neko.GetText())
			if message.From.ID == message.Chat.ID {
				msg.ReplyToMessageID = message.MessageID
			}
		} else {
			msg = tg.NewMessage(message.Chat.ID, neko.GetText())
			if reply && message.ReplyToMessage != nil {
				msg.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				msg.ReplyToMessageID = message.MessageID
			}
		}

		if hasMsg {
			msg.Text += text
		}
		sendIt = msg
	} else if neko.IsVideo() {
		var video tg.VideoConfig
		if pv {
			video = tg.NewVideo(message.From.ID, tg.FileURL(neko.GetUrl()))
			if message.From.ID == message.Chat.ID {
				video.ReplyToMessageID = message.MessageID
			}
		} else {
			video = tg.NewVideo(message.Chat.ID, tg.FileURL(neko.GetUrl()))
			if reply && message.ReplyToMessage != nil {
				video.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				video.ReplyToMessageID = message.MessageID
			}
		}
		if hasMsg {
			video.Caption = text
		}
		sendIt = video
	} else {
		log.Println(neko.GetUrl())
		return
	}

	if del {
		req := tg.NewDeleteMessage(message.Chat.ID, message.MessageID)
		// don't check error or response, we have
		// more important things to do
		go settings.GetAPI().Request(req)
	}

	if _, err := api.Send(sendIt); err != nil {
		log.Println("err in sending ", err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
	}
}

func getCorrectText(message *tg.Message, args *pTools.EventArgs) string {
	var str string
	isRelpy := message.ReplyToMessage != nil
	if isRelpy {
		isRelpy = !wotoStrings.IsEmpty(&message.ReplyToMessage.Text)
		if !isRelpy {
			isRelpy = !wotoStrings.IsEmpty(&message.ReplyToMessage.Caption)
			if !isRelpy {
				str = args.GetAsStringOrRaw(TEXT_FLAG, TXT_FLAG)
			}
		} else {
			str = message.ReplyToMessage.Text
		}
	} else {
		str = args.GetAsStringOrRaw(TEXT_FLAG, TXT_FLAG)
	}

	if wotoStrings.IsEmpty(&str) {
		return wotoValues.EMPTY
	}

	return str
}
