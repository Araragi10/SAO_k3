// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import (
	"errors"
	"log"
	"reflect"

	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
)

// GetHPatList will give you the pat list in the db.
// Even if there is no pat in the db, `[]bson.M` will not be nil,
// though its length will be zero.
func (_c *WotoClient) GetPatList() ([]bson.M, error) {
	if !_c.isPatClient() {
		return nil, errors.New("you can't use this client for pat plugin")
	}

	raws, err := _c.getRAWS(PatDataBase, PatListCollection)
	if err != nil {
		return nil, err
	}

	if raws == nil {
		return nil, errors.New("an unknown error in get pat list")
	}

	return raws, nil
}

// GetHPatList will give you the hentai pat list in the db.
// Even if there is no pat in the db, `[]bson.M` will not be nil,
// though its length will be zero.
func (_c *WotoClient) GetHPatList() ([]bson.M, error) {
	if !_c.isPatClient() {
		return nil, errors.New("wrong usage of wrong wclient," +
			"this client is not a pat client")
	}

	raws, err := _c.getRAWS(PatDataBase, HPatListCollection)
	if err != nil {
		return nil, err
	}

	if raws == nil {
		return nil, errors.New("an unknown error occured in getting" +
			"raw data from db, because it's nil")
	}
	return raws, nil
}

// AddPat will add a new pat to the db.
// it's not our duty to check if the patID already exists or not!
// pat plugin should check if before call this method!
func (_c *WotoClient) AddPat(patID string, t int32) ([]bson.M, error) {
	if !_c.isPatClient() {
		return nil, errors.New("wrong usage of non-pat wclient in " +
			"add pat function")
	}

	collection := _c.getCollection(PatDataBase, PatListCollection)

	b := bson.D{
		{Key: wv.PAT_ID_KEY, Value: patID},
		{Key: wv.PAT_TYPE_KEY, Value: t},
	}
	_, err := collection.InsertOne(*_c.ctx, b)

	if err != nil {
		return nil, err
	}

	return _c.GetPatList()
}

// AddPat will add a new hentai pat to the db.
// it's not our duty to check if the patID already exists or not!
// pat plugin should check if before call this method!
func (_c *WotoClient) AddHPat(patID string, t int32) ([]bson.M, error) {
	if !_c.isPatClient() {
		return nil, errors.New("wrong usage of non-pat wclient " +
			"in add hentai pat function")
	}

	collection := _c.getCollection(PatDataBase, HPatListCollection)

	b := bson.D{
		{Key: wv.PAT_ID_KEY, Value: patID},
		{Key: wv.PAT_TYPE_KEY, Value: t},
	}
	_, err := collection.InsertOne(*_c.ctx, b)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return _c.GetHPatList()
}

// RemovePat will remove the pat with specified id.
// It will return `FAILED` if it encouners any error,
// will return `SUCCESS` if it successfully deleted the pat and
// the call to getPatList was successful,
// and will return `CANCELED` if it found nothing in loop.
// It's not that it will give you information about existing of a pat!
// so be carefull about it; the pat MAY EXISTS in the db, but
// the method may return `CANCELED`.
func (_c *WotoClient) RemovePat(patID string) ([]bson.M, error) {
	if !_c.isPatClient() {
		return nil, errors.New("wrong usage of non-pat wclient in " +
			"remove pat function")
	}

	collection := _c.getCollection(PatDataBase, PatListCollection)
	raws, err := _c.getRAWS(PatDataBase, PatListCollection)
	if err != nil {
		return nil, err
	}

	if raws == nil {
		return nil, errors.New("got nil raws data in remove pat function")
	}

	id := wv.EMPTY
	for _, current := range raws {
		if reflect.TypeOf(current[wv.PAT_ID_KEY]) != reflect.TypeOf(id) {
			continue
		}

		id = current[wv.PAT_ID_KEY].(string)
		if id == patID {
			collection.DeleteOne(*_c.ctx, current)
			return _c.GetPatList()
		}
	}

	return nil, errors.New("pat doesn't exist in the db at all")
}

// RemoveHPat will remove the hentai pat with specified id.
// It will return `FAILED` if it encouners any error,
// will return `SUCCESS` if it successfully deleted the pat and
// the call to getPatList was successful,
// and will return `CANCELED` if it found nothing in loop.
// It's not that it will give you information about existing of a pat!
// so be carefull about it; the pat MAY EXISTS in the db, but
// the method may return `CANCELED`.
func (_c *WotoClient) RemoveHPat(patID string) ([]bson.M, error) {
	if !_c.isPatClient() {
		return nil, errors.New("wrong usage")
	}

	collection := _c.getCollection(PatDataBase, HPatListCollection)
	raws, err := _c.getRAWS(PatDataBase, HPatListCollection)
	if err != nil {
		return nil, err
	}

	id := wv.EMPTY
	for _, current := range raws {
		if reflect.TypeOf(current[wv.PAT_ID_KEY]) != reflect.TypeOf(id) {
			continue
		}
		id = current[wv.PAT_ID_KEY].(string)
		if id == patID {
			collection.DeleteOne(*_c.ctx, current)
			return _c.GetPatList()
		}
	}

	return raws, nil
}
