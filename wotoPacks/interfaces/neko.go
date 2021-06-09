// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

type NekoFactor interface {
	IsPhoto() bool
	IsGif() bool
	IsText() bool
	IsVideo() bool
	GetText() string
	GetUrl() string
}
