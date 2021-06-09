// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SudoInfo interface {
	GetID() int64
	GetNickname() string
	GetDate() time.Time
	GetAsP() *primitive.M
}

type SudoList interface {
	GetSudo(id *int64) SudoInfo
	Contains(id *int64) bool
	GetListP() []primitive.M
}
