// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"context"
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gotd/td/tg"
)

func (_s *AppSettings) SetAPI(_api *tgbot.BotAPI) {
	_s.botAPI = _api
}

func (_s *AppSettings) SetIndex(_index int) {
	_s.index = _index
}

// SetSudoList will set the sudo list.
// sudo_id = the user id of sudo. (int64)
// sudo_date = the date which it becomes sudo. (see time)
// sudo_nickname = the nickname of sudo. (string)
func (_s *AppSettings) SetSudoList(list interfaces.SudoList) {
	_s.sudoList = list
}

func (_s *AppSettings) SetMainSudo(id int64) {
	if _s.mainSudo != id {
		_s.mainSudo = id
	}
}

func (_s *AppSettings) SetTObt(_obt string) {
	_s.tObt = _obt
}

func (_s *AppSettings) AddSudo(id int64, nick string) wa.RESULT {
	if !_s.IsSudo(id) {
		re := _s.wClient.AddSudo(id, nick)
		if re != wa.SUCCESS {
			log.Println(wv.SUDO_ADD_SETTINGS)
		}
		return re
	} else {
		return wa.CANCELED
	}
}

func (_s *AppSettings) RemSudo(id int64) wa.RESULT {
	if _s.IsSudo(id) {
		re, err := _s.wClient.RemSudo(id)
		if re != wa.SUCCESS {
			log.Println(err)
		}
		return re
	} else {
		return wa.CANCELED
	}
}

func (_s *AppSettings) SetWClient(_client interfaces.WClient) {
	_s.wClient = _client
}

func (_s *AppSettings) SetGClient(g *tg.Client, ctx *context.Context) {
	_s.gClient = g
	_s.gCtx = ctx
}
