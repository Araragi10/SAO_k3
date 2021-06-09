// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"context"
	"time"

	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gotd/td/tg"
)

var _alreadyRunning bool
var _settings *AppSettings

type AppSettings struct {
	index      int
	totalIndex int
	tObt       string
	isGlobal   bool
	mainSudo   int64
	wClient    interfaces.WClient
	patClient  interfaces.WClient
	botAPI     *tgbotapi.BotAPI
	gClient    *tg.Client
	gCtx       *context.Context
	started    *time.Time
	sudoList   interfaces.SudoList
}

func GetSettings() interfaces.WSettings {
	if _settings != nil {
		return GetExisting()
	}
	settings := AppSettings{}
	rn := time.Now()
	settings.started = &rn
	_settings = &settings
	return _settings
}

func GetExisting() interfaces.WSettings {
	return _settings
}

func App_enter() {
	_alreadyRunning = true
}

func App_exit() {
	_alreadyRunning = false
}

func IsRunning() bool {
	return _alreadyRunning
}
