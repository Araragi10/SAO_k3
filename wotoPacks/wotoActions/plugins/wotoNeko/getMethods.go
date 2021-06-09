// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

import (
	"math/rand"
	"strings"

	ws "github.com/Araragi10/SAO_k3/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
)

//---------------------------------------------------------

func (n *NekoBase) IsPhoto() bool {
	tmp := strings.ToLower(n.Url)
	b1 := strings.HasSuffix(tmp, wv.Jpg)
	b2 := strings.HasSuffix(tmp, wv.Jpeg)
	b3 := strings.HasSuffix(tmp, wv.Png)
	return b1 || b2 || b3
}

func (n *NekoBase) IsGif() bool {
	tmp := strings.ToLower(n.Url)
	return strings.HasSuffix(tmp, wv.Gif)
}

func (n *NekoBase) IsVideo() bool {
	tmp := strings.ToLower(n.Url)
	return strings.HasSuffix(tmp, wv.Mp4)
}

func (n *NekoBase) IsText() bool {
	b1 := ws.IsEmpty(&n.Url)
	b2 := !ws.IsEmpty(&n.Why)
	b3 := !ws.IsEmpty(&n.Name)
	b4 := !ws.IsEmpty(&n.Cat)
	b5 := !ws.IsEmpty(&n.Fact)
	b6 := !ws.IsEmpty(&n.Owo)
	return b1 && (b2 || b3 || b4 || b5 || b6)
}

func (n *NekoBase) GetText() string {
	if !n.IsText() {
		return wv.EMPTY
	}

	if !ws.IsEmpty(&n.Why) {
		return n.Why
	}
	if !ws.IsEmpty(&n.Name) {
		return n.Name
	}
	if !ws.IsEmpty(&n.Cat) {
		return n.Cat
	}
	if !ws.IsEmpty(&n.Fact) {
		return n.Fact
	}
	if !ws.IsEmpty(&n.Owo) {
		return n.Owo
	}

	return wv.EMPTY
}

func (n *NekoBase) GetUrl() string {
	return n.Url
}

//---------------------------------------------------------

func (n *TenorBase) IsPhoto() bool {
	tmp := strings.ToLower(n.Src)
	b1 := strings.HasSuffix(tmp, wv.Jpg)
	b2 := strings.HasSuffix(tmp, wv.Jpeg)
	b3 := strings.HasSuffix(tmp, wv.Png)
	return b1 || b2 || b3
}

func (n *TenorBase) IsGif() bool {
	tmp := strings.ToLower(n.Src)
	if strings.HasSuffix(tmp, wv.Gif) {
		return true
	}

	switch n.MediaType {
	case string(tMediumgifType),
		string(tNanogifType),
		string(tGifType),
		string(tTinygifType):
		return true
	}

	return false
}

func (n *TenorBase) IsVideo() bool {
	tmp := strings.ToLower(n.Src)
	if strings.HasSuffix(tmp, wv.Gif) {
		return true
	}

	switch n.MediaType {
	case string(tMp4Type),
		string(tTinywebmType),
		string(tTinymp4Type),
		string(tNanowebmType),
		string(tNanomp4Type),
		string(tLoopedmp4Type),
		string(tWebmType):
		return true
	}

	return false
}

func (n *TenorBase) IsText() bool {
	return false
}

func (n *TenorBase) GetText() string {
	return wv.EMPTY
}

func (n *TenorBase) GetUrl() string {
	return n.Src
}

// GetPreview will give you the preview webpage
// link on tenor website.
func (n *TenorBase) GetPreview() string {
	return n.Url
}

//---------------------------------------------------------

func (n *TenorList) GetRandom() *TenorBase {
	l := len(*n)
	if l == wv.BaseIndex {
		return nil
	}

	return &(*n)[rand.Intn(l)]
}

func (n *TenorList) IsEmpty() bool {
	return len(*n) == wv.BaseIndex
}

//---------------------------------------------------------

func (c *TenorContainer) IsEmpty() bool {
	return len(c.Results) == wv.BaseIndex
}

// GetRandom will give you a random result from a
// tenor container.
func (c *TenorContainer) GetRandom() *ResultsTenor {
	l := len(c.Results)
	if l == wv.BaseIndex {
		return nil
	}

	return &c.Results[rand.Intn(l)]
}

//---------------------------------------------------------

func (r *ResultsTenor) IsEmptyMedia() bool {
	return r.MediaLength() == wv.BaseIndex
}

func (r *ResultsTenor) MediaLength() int {
	return len(r.Media)
}

func (r *ResultsTenor) GetRandomMedia() *MediaTenor {
	if r.IsEmptyMedia() {
		return nil
	}

	return &r.Media[rand.Intn(r.MediaLength())]
}

//---------------------------------------------------------
