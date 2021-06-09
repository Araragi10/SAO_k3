// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package textMessage

import (
	"github.com/Araragi10/SAO_k3/wotoPacks/appSettings"
	bc "github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/messages/textMessage/botCommands"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleTextMessage(message *tg.Message) {
	if message == nil {
		return
	}

	if bc.IsSudoCommand(&message.Text) {
		s := appSettings.GetExisting()
		if s == nil {
			return
		}

		id := message.From.ID
		if s.IsSudo(id) || s.IsMainSudo(id) {
			bc.HandleSudoCommand(message)
		}
		return
	} else if bc.IsCommand(&message.Text) {
		//log.Println("Is HERE!")
		bc.HandleCommand(message)
	} //else {
	//log.Println("In ELSeE!")
	//}

}
