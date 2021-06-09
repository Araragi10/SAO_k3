package wotoTest

import (
	"log"

	"github.com/Araragi10/SAO_k3/wotoPacks/appSettings"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/pTools"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TestCommandHandle is not completed yet!
func TestCommandHandle(message *tg.Message, args *pTools.EventArgs) {
	settings := appSettings.GetExisting()
	//log.Println("In EVENT!")
	if settings == nil {
		return
	}
	var str string
	if args.IsEmptyOrRaw() {
		str = "no test session provided!"
	} else {
		str = "couldn't find test session : " +
			args.GetAsStringOrRaw(SESSION_FLAG, SE_FLAG)
	}
	msg := tg.NewMessage(message.Chat.ID, str)
	msg.ReplyToMessageID = message.MessageID
	if _, err := settings.GetAPI().Send(msg); err != nil {
		log.Println(err)
		return
	}
}
