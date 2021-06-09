// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/pTools"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/wotoGP"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/wotoMorse"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/wotoNeko"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/wotoPat"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/wotoTranslate"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/wotoUD"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func toMorseHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoMorse.ToMorseHandler(message, args)
}

func fromMorseHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoMorse.FromMorseHandler(message, args)
}

func trHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoTranslate.TrHandler(message, args)
}

func patHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoPat.PatHandler(message, args)
}

func tickleNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.TickleHandler(message, args)
}

func slapNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.SlapHandler(message, args)
}

func pokeNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.PokeHandler(message, args)
}

func nekoNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.NekoHandler(message, args)
}

func meowNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.MeowHandler(message, args)
}

func lizardNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.LizardHandler(message, args)
}

func kissNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.KissHandler(message, args)
}

func hugNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.HugHandler(message, args)
}

func foxNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.FoxHandler(message, args)
}

func feedNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.FeedHandler(message, args)
}

func cuddleNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.CuddleHandler(message, args)
}

func kemonomimiNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.KemonomimiHandler(message, args)
}

func holoNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.HoloHandler(message, args)
}

func smugNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.SmugHandler(message, args)
}

func bakaNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.BakaHandler(message, args)
}

func woofNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.WoofHandler(message, args)
}

func gooseNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.GooseHandler(message, args)
}

func gecgNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.GecgHandler(message, args)
}

func avatarNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.AvatarHandler(message, args)
}

func waifuNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.WaifuHandler(message, args)
}

func whyNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.WhyHandler(message, args)
}

func nameNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.NameHandler(message, args)
}

func catNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.CatHandler(message, args)
}

func factNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.FactHandler(message, args)
}

func owoNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.OwoHandler(message, args)
}

func gifDNekoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.TGifDHandler(message, args)
}

func mp4Handler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.Mp4Handler(message, args)
}

func mediumgifHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.MediumgifHandler(message, args)
}

func nanogifHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.NanogifHandler(message, args)
}

func tinywebmHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.TinywebmHandler(message, args)
}

func gifHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.GifHandler(message, args)
}

func tinymp4Handler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.Tinymp4Handler(message, args)
}

func nanowebmHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.NanowebmHandler(message, args)
}

func nanomp4Handler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.Nanomp4Handler(message, args)
}

func loopedmp4Handler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.Loopedmp4Handler(message, args)
}

func webmHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.WebmHandler(message, args)
}

func tinygifHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoNeko.TinygifHandler(message, args)
}

func udHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoUD.UdHandler(message, args)
}

func pingHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoGP.PingHandler(message, args)
}

func banHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoGP.BanHandler(message, args)
}

func infoHandler(message *tg.Message, args *pTools.EventArgs) {
	wotoGP.InfoHandler(message, args)
}
