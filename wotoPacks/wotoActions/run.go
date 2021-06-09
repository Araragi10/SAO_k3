// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"log"

	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunBot(settings interfaces.WSettings) {
	bot, err := tg.NewBotAPI(settings.GetObt())
	settings.SetAPI(bot)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Debug = false

	settings.GetWClient().PingClientDB(true)

	//tClient := settings.GetTDLibClient()

	//u, err := tClient.GetMe()

	//u, err := tClient.SearchChats("@Falling_inside_The_Black", 1)
	//u, err := tClient.SearchContacts("Falling_inside_The_Black", 100)
	//if err != nil {
	//	settings.SendSudo("FINDING ERR: " + err.Error())
	//}

	//d, err := json.MarshalIndent(u, "", "  ")
	//if err != nil {
	//	settings.SendSudo("in JSON ERR: " + err.Error())
	//}

	//str := "here we go\n" + string(d)
	//	for _, current := range u.Bio {
	//		str += "ID: " + strconv.Itoa(int(current)) + "\n"
	//
	//}
	//settings.SendSudo(str)

	for {
		runOnce(bot, settings)
	}
}

func runOnce(_bot *tg.BotAPI, _settings interfaces.WSettings) {
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tg.NewUpdate(wotoValues.BaseIndex)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = wotoValues.BaseTimeOut

	// Start polling Telegram for updates.
	updates := _bot.GetUpdatesChan(updateConfig)

	// go through each update that we're getting from Telegram.
	for update := range updates {
		// check if the current application is allowed to handle the update
		// request or not.
		if !shouldHandle(&update) {
			continue
		}
		go HandleMessage(&update, _settings)

	}
}
