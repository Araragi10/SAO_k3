// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

import (
	wa "github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/common"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoDB/dbTypes"
	"go.mongodb.org/mongo-driver/bson"
)

type WClient interface {
	PingClientDB(_report bool)
	DeleteCollection(database dbTypes.DATABASE, collection dbTypes.COLLECTION) wa.RESULT
	Destroy()
	GetWotoConfiguration() (wa.RESULT, error)
	GetPatList() ([]bson.M, error)
	GetHPatList() ([]bson.M, error)
	ResetWotoConfiguration() (wa.RESULT, error)
	ResetUsersCollection() wa.RESULT
	CreateNewConfiguration() wa.RESULT
	AddSudo(id int64, nick string) wa.RESULT
	RemSudo(id int64) (wa.RESULT, error)
	AddPat(patID string, t int32) ([]bson.M, error)
	AddHPat(patID string, t int32) ([]bson.M, error)
	RemovePat(patID string) ([]bson.M, error)
	RemoveHPat(patID string) ([]bson.M, error)
	FindAccount(_username, _pass *string) wa.RESULT
	DeleteAccount() wa.RESULT
	UpdateAccount() wa.RESULT
	UpdateAccountOnlineToken() wa.RESULT
}
