// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TgError interface {
	SendRandomErrorMessage(*tgbotapi.Message)
}
