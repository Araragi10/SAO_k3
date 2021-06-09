// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

type GbanInfo interface {
	GetTarget() int64
	GetSudo() SudoInfo
	GetReason() string
	GetMessageID() int // message id in channel.
}
