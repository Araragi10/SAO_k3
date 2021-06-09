// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"log"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IsSudoCommand(_text *string) bool {
	return strings.HasPrefix(*_text, wotoValues.SUDO_PREFIX1)
}

func HandleSudoCommand(message *tgbotapi.Message) {
	sudoListInit()

	eventArg, err := pTools.ParseArg(message.Text)
	if err != nil {
		log.Println("an error in parsing the args:", err)
		return
	}

	if eventArg == nil {
		log.Println("unexpected error, event arg is nil")
		return
	}

	event, ok := sudoCMDList[eventArg.GetCommand()]
	if event != nil && ok {
		event(message, eventArg)
	}
}
