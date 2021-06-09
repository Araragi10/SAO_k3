// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import (
	"errors"

	wa "github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/common"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/wotoSudo"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoDB/dbTypes"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
)

// GetWotoConfiguration will directly give you the configuration from db.
func (_c *WotoClient) GetWotoConfiguration() (result wa.RESULT, err error) {
	raws, err := _c.getRAWS(MainDataBase, ConfigurationCollection)
	if err != nil {
		return wa.FAILED, err
	}

	_c.settings.SetMainSudo(wv.SUDO)
	_c.settings.SetSudoList(wotoSudo.ToSudoList(raws))

	//appSettings.GetExisting().SendSudo(strconv.FormatInt(mainSudo, wotoValues.BaseTen))

	//var currentValue string
	//for _, current := range raws {
	//	currentValue := current["AdminList"].(string)
	//	if
	//}
	//str := raws[0]["AdminList"].(string)
	//appSettings.GetExisting().SendSudo(str)
	//_w := wotoConfiguration{
	//	UIDKeyName: [MaxUIDIndex]string{},
	//	LastUID:    [MaxUIDIndex]UID{},
	//}
	//_re, _raws := _c.getRAWS(MainDataBase, ConfigurationCollection)
	//if _re != SUCCESS {
	//	return FAILED, nil
	//}
	//_w.setUIDKeys()
	//for _i, _current := range _raws {
	//	_value, _ok := _current[UIDKeyName].(string)
	//	_w.LastUID[_i] = UID(_value)
	//	if !_ok {
	//		clientError()
	//		return FAILED, nil
	//	}
	//}
	return wa.SUCCESS, nil
}

// getRAWS will give you the raw data.
func (_c *WotoClient) getRAWS(database dbTypes.DATABASE,
	collection dbTypes.COLLECTION) (m []bson.M, err error) {
	_collection := _c.getCollection(database, collection)
	if _collection == nil {
		// clientError()
		// well, it seems it's really rare to reach this poit,
		// but I will note this so do something about it in the future.
		return nil, errors.New("got a nil collection from db")
	}
	cursor, cursorError := _collection.Find(*_c.ctx, bson.M{})
	if cursorError != nil {
		return nil, cursorError
	}
	if cursor == nil {
		return nil, errors.New("got a nil cursor from db")
	}

	var raws []bson.M
	err = cursor.All(*_c.ctx, &raws)
	if err != nil {
		return nil, err
	}

	return raws, nil
}

func (_c *WotoClient) isRawConfigClient() bool {
	return _c.index == RAW_CONFIG_CLIENT
}

func (_c *WotoClient) isPatClient() bool {
	return _c.index == PAT_CLIENT
}
