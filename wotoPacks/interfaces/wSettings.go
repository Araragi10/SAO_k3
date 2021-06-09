// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

import (
	"context"

	wa "github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/common"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gotd/td/tg"
)

// Go encourages composition over inheritance,
// using simple, often one-method interfaces ...
// that serve as clean, comprehensible boundaries between components.
// 	—Rob Pike,
//		“Go at Google: Language Design in the
//		Service of Software Engineering”
//		(see talks.golang.org/ 2012/splash.article)

// WSettings is the settings interface for application
type WSettings interface {
	GetObt() string
	GetAPI() *tgbot.BotAPI
	IsGlobal() bool
	GetSudoList() SudoList
	IsSudo(id int64) bool
	IsMainSudo(id int64) bool
	GetWClient() WClient
	GetPatClient() WClient
	GetGClient() (*tg.Client, *context.Context)
	SetAPI(_api *tgbot.BotAPI)
	SetSudoList(list SudoList)
	SetTObt(_obt string)
	SetMainSudo(id int64)
	AddSudo(id int64, nick string) wa.RESULT
	RemSudo(id int64) wa.RESULT
	SendSudo(str string)
	// set the main w client.
	SetWClient(_client WClient)
	// set the pat db client.
	SetPatClient(client WClient)
	// set the Gclient. a client for do special actions
	// using mtproto.
	// like getting information of a user by its username.
	SetGClient(g *tg.Client, ctx *context.Context)
}
