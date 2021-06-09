// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var tenorMap map[int64]bool
var tenorMutix *sync.RWMutex

// TenorSearch performs a search operations in tenor.
// it will return you a `*TenorContainer`, and in a case
// that there is an error, it will set the err as well.
func TenorSearch(que string) (container *TenorContainer, err error) {
	target := tContainerHostUrl + url.QueryEscape(que) +
		tContainerKeyUrl + gOuTeMceept +
		tContainerLimitUrl + defaultLimit

	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &container)
	if err != nil {
		return nil, err
	}

	return
}

func GetTenorRandomMedia(que string) (media *MediaTenor, err error) {
	container, err := TenorSearch(que)
	if err != nil {
		return nil, err
	}

	if container.IsEmpty() {
		return nil, errors.New("not found")
	}

	r := container.GetRandom()
	if r == nil {
		return nil, errors.New("got a nil result from container")
	}

	return r.GetRandomMedia(), nil
}

func GetTenorRandomMp4(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	mp4 := m.Mp4
	if mp4 == nil {
		return nil, errors.New("no mp4 version is available")
	}

	return toTenorBase(&mp4.URL, &mp4.Preview, tMp4Type)
}

func GetTenorRandomMediumgif(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	mg := m.Mediumgif
	if mg == nil {
		return nil, errors.New("no medium gif version is available")
	}

	return toTenorBase(&mg.URL, &mg.Preview, tMediumgifType)
}

func GetTenorRandomNanogif(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	ng := m.Nanogif
	if ng == nil {
		return nil, errors.New("no nano gif version is available")
	}

	return toTenorBase(&ng.URL, &ng.Preview, tNanogifType)
}

func GetTenorRandomTinywebm(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	wm := m.Tinywebm
	if wm == nil {
		return nil, errors.New("no tiny webm version is available")
	}

	return toTenorBase(&wm.URL, &wm.Preview, tWebmType)
}

func GetTenorRandomGif(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	gif := m.Gif
	if gif == nil {
		return nil, errors.New("no gif version is available")
	}

	return toTenorBase(&gif.URL, &gif.Preview, tGifType)
}

func GetTenorRandomTinymp4(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	mp4 := m.Tinymp4
	if mp4 == nil {
		return nil, errors.New("no tiny mp4 version is available")
	}

	return toTenorBase(&mp4.URL, &mp4.Preview, tTinymp4Type)
}

func GetTenorRandomNanowebm(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	nw := m.Nanowebm
	if nw == nil {
		return nil, errors.New("no nano webm version is available")
	}

	return toTenorBase(&nw.URL, &nw.Preview, tNanowebmType)
}

func GetTenorRandomNanomp4(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	mp4 := m.Nanomp4
	if mp4 == nil {
		return nil, errors.New("no nano mp4 version is available")
	}

	return toTenorBase(&mp4.URL, &mp4.Preview, tNanomp4Type)
}

func GetTenorRandomLoopedmp4(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	lm := m.Loopedmp4
	if lm == nil {
		return nil, errors.New("no looped mp4 version is available")
	}

	return toTenorBase(&lm.URL, &lm.Preview, tLoopedmp4Type)
}

func GetTenorRandomWebm(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	wm := m.Webm
	if wm == nil {
		return nil, errors.New("no webm version is available")
	}

	return toTenorBase(&wm.URL, &wm.Preview, tWebmType)
}

func GetTenorRandomTinygif(que string) (base *TenorBase, err error) {
	m, err := GetTenorRandomMedia(que)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("got a nil random media")
	}

	tg := m.Tinygif
	if tg == nil {
		return nil, errors.New("no tiny gif version is available")
	}

	return toTenorBase(&tg.URL, &tg.Preview, tTinygifType)
}

func checkTenor(message *tgbotapi.Message) (send *tgbotapi.Message, ok bool) {
	if tenorMap == nil {
		tenorMap = make(map[int64]bool)
	}

	if tenorMutix == nil {
		tenorMutix = &sync.RWMutex{}
	}

	tenorMutix.Lock()

	if tenorMap[message.Chat.ID] {
		tenorMutix.Unlock()
		return nil, false
	} else {
		tenorMap[message.Chat.ID] = true
	}

	tenorMutix.Unlock()

	s := appSettings.GetExisting()
	if s == nil {
		return nil, false
	}

	api := s.GetAPI()
	if api == nil {
		return nil, false
	}

	md := wotoMD.GetItalic(searchingMessage)
	if md == nil {
		return nil, false
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		md.ToString())

	msg.ParseMode = tgbotapi.ModeMarkdownV2
	msg.ReplyToMessageID = message.MessageID

	resp, err := api.Send(msg)
	if err != nil {
		log.Println(err)
	}

	return &resp, true
}

func endTenor(message *tgbotapi.Message) {
	if tenorMap == nil {
		tenorMap = make(map[int64]bool)
	}

	tenorMutix.Lock()
	delete(tenorMap, message.Chat.ID)
	tenorMutix.Unlock()

	s := appSettings.GetExisting()
	if s == nil {
		return
	}

	api := s.GetAPI()
	if api == nil {
		return
	}

	dl := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)

	api.Request(dl)
}
