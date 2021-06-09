package wotoDB

/*
import (
	"errors"
	"log"
	"reflect"

	wa "github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/common"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
)

// GetHPatList will give you the pat list in the db.
// Even if there is no pat in the db, `[]bson.M` will not be nil,
// though its length will be zero.
func (_c *WotoClient) GetPatList2() (wa.RESULT, []bson.M) {
	if !_c.isRawConfigClient() {
		return wa.FAILED, nil
	}

	re, raws := _c.getRAWS(PatDataBase, PatListCollection)
	if re != wa.SUCCESS || raws == nil {
		log.Println(wv.GET_PAT_LIST_ERROR)
		return wa.FAILED, nil
	}
	return wa.SUCCESS, raws
}

// AddPat will add a new pat to the db.
// it's not our duty to check if the patID already exists or not!
// pat plugin should check if before call this method!
func (_c *WotoClient) AddPat2(patID string, t int32) (wa.RESULT, []bson.M) {
	if !_c.isRawConfigClient() {
		return wa.FAILED, nil
	}

	collection := _c.getCollection(PatDataBase, PatListCollection)

	b := bson.D{
		{Key: wv.PAT_ID_KEY, Value: patID},
		{Key: wv.PAT_TYPE_KEY, Value: t},
	}
	_, err := collection.InsertOne(*_c.ctx, b)

	if err != nil {
		log.Println(err)
		return wa.FAILED, nil
	}

	return _c.GetPatList()
}

// RemovePat will remove the pat with specified id.
// It will return `FAILED` if it encouners any error,
// will return `SUCCESS` if it successfully deleted the pat and
// the call to getPatList was successful,
// and will return `CANCELED` if it found nothing in loop.
// It's not that it will give you information about existing of a pat!
// so be carefull about it; the pat MAY EXISTS in the db, but
// the method may return `CANCELED`.
func (_c *WotoClient) RemovePat2(patID string) ([]bson.M, error) {
	if !_c.isRawConfigClient() {
		return nil, errors.New("wrong usage")
	}

	collection := _c.getCollection(PatDataBase, PatListCollection)
	raws, err := _c.getRAWS(PatDataBase, PatListCollection)
	if err != nil {
		return nil, err
	}

	if raws == nil {
		return nil, errors.New("got nil raw data from db")
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

*/
