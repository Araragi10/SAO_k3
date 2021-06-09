// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/messages/callBackQ"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/messages/textMessage"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var updatesMap map[int]bool

// shouldHangle will check if you should handle the
// update or not.
func shouldHandle(update *tg.Update) bool {
	return true
}

// HandleMessage will handle the update comming from the telegram servers.
func HandleMessage(update *tg.Update, settings interfaces.WSettings) {
	switch getMessageType(update) {
	case NONE:
		return
	case CALLBACK_QUERY:
		callBackQ.QHandler(update.CallbackQuery)
	case TEXT_MESSAGE:
		textMessage.HandleTextMessage(update.Message)
	default:
		return
	}
}
