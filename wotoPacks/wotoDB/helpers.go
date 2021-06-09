// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import (
	"errors"
	"log"
	"time"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
)

func (_c *WotoClient) CreateNewConfiguration() wa.RESULT {
	collection := _c.getCollection(MainDataBase, ConfigurationCollection)
	//SudoList := strconv.FormatInt(wv.SUDO, wv.BaseTen)
	b := bson.D{
		{Key: wv.SudoID, Value: wv.SUDO},
		{Key: wv.SudoNick, Value: wv.AutoStr},
		{Key: wv.SudoDate, Value: time.Now()},
	}

	_, err := collection.InsertOne(*_c.ctx, b)
	if err != nil {
		appSettings.GetExisting().SendSudo(err.Error())
	}

	return wa.SUCCESS
}

func (_c *WotoClient) AddSudo(id int64, nick string) wa.RESULT {
	if id == wv.BaseIndex {
		log.Println("The ID cannot be zero!")
		return wa.FAILED
	}

	if _c.settings.IsSudo(id) {
		log.Println("This user is already a sudo!")
		return wa.CANCELED
	}

	collection := _c.getCollection(MainDataBase, ConfigurationCollection)

	b := bson.D{
		{Key: wv.SudoID, Value: id},
		{Key: wv.SudoNick, Value: nick},
		{Key: wv.SudoDate, Value: time.Now()},
	}

	_, err := collection.InsertOne(*_c.ctx, b)
	if err != nil {
		appSettings.GetExisting().SendSudo(err.Error())
	}

	_c.GetWotoConfiguration()

	return wa.SUCCESS
}

func (_c *WotoClient) RemSudo(id int64) (wa.RESULT, error) {

	if id == wv.BaseIndex {
		return wa.FAILED, errors.New("user id cannot be zero")
	}

	if !_c.settings.IsSudo(id) {
		return wa.CANCELED, errors.New("this user is not a sudo at all")
	}
	if _c.settings.IsMainSudo(id) {
		return wa.CANCELED, errors.New("this use is my main sudo, " +
			"you shall not remove this person")
	}

	collection := _c.getCollection(MainDataBase, ConfigurationCollection)

	list := appSettings.GetExisting().GetSudoList()

	su := list.GetSudo(&id)
	if su == nil {
		return wa.FAILED, errors.New("there is a problem in getting " +
			"info of this user")
	}

	cL := su.GetAsP()
	if cL == nil {
		return wa.FAILED, errors.New("tried to get primitive value of " +
			"this user, but it failed")
	}

	_, err := collection.DeleteOne(*_c.ctx, cL)
	if err != nil {
		return wa.FAILED, err
	}

	return _c.GetWotoConfiguration()
}

func (_c *WotoClient) AddGBan(info interfaces.GbanInfo) error {
	if info.GetTarget() == wv.BaseIndex {
		log.Println("The ID cannot be zero!")
		return errors.New("target id should not be zero")
	}

	if _c.settings.IsSudo(info.GetTarget()) {
		return errors.New("cannot gban a sudo")
	}

	sudo := info.GetSudo()
	if sudo == nil {
		return errors.New("sudo request sender cannot be nil")
	}

	collection := _c.getCollection(MainDataBase, gBanUsersCollection)

	b := bson.D{
		{Key: wv.SudoID, Value: info.GetTarget()},
		{Key: wv.SudoNick, Value: info.GetSudo().GetNickname()},
		{Key: wv.SudoDate, Value: time.Now()},
	}

	_, err := collection.InsertOne(*_c.ctx, b)
	if err != nil {
		appSettings.GetExisting().SendSudo(err.Error())
	}

	_c.GetWotoConfiguration()

	return nil
}

func (_c *WotoClient) RemGBan(id int64, sudo int64, nick string) wa.RESULT {
	return wa.FAILED
}
