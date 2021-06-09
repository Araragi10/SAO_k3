package wotoNeko

import (
	"errors"
	"net/url"

	ws "github.com/Araragi10/SAO_k3/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
)

func genDAuth() string {
	return dAuth01 +
		wv.Point + dAuth02 +
		wv.Point + dAuth03
}

func getDreq(q *string) string {

	return tGifHostUrl + url.QueryEscape(*q)
}

// toTenorBase will give you a tenor base
func toTenorBase(rlSrc, pre *string, t tenorType) (base *TenorBase, err error) {
	if ws.IsEmpty(rlSrc) {
		return nil, errors.New("real source url string cannot be empty")
	}

	return &TenorBase{
		Url:       *pre,
		Src:       *rlSrc,
		MediaType: string(t),
	}, nil
}
