package wotoTranslate

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrong"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

//import "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoTranslate/wotoLang"

func DetectLanguage(text string) *Lang {
	m := map[string]string{
		userAgentKey:      userAgentValue,
		acceptKey:         acceptValue,
		acceptLanguageKey: acceptLanguageValue,
		refererKey:        refererValue,
		contentTypeKey:    contentTypeValue,
		originKey:         originValue,
		connectionKey:     connectionValue,
		teKey:             teValue,
		qKey:              text,
	}

	data, errJ := json.Marshal(m)
	if errJ != nil {
		log.Println(errJ)
		return nil
	}

	reader := bytes.NewReader(data)
	resp, errH := http.Post(dHostUrl, contentTypeValue, reader)

	if errH != nil {
		log.Println(errH)
	}

	defer resp.Body.Close()

	b, errB := ioutil.ReadAll(resp.Body)
	if errB != nil {
		log.Println(errB)
	}

	log.Println(string(b))
	str := ws.Qsb(b)
	strs := str.SplitStr(preLeft, preRight)
	if len(strs) <= wv.BaseOneIndex {
		for _, tr := range strs {
			log.Println("HERE: " + tr.GetValue())
		}
		return nil
	}

	b = []byte(strs[wv.BaseOneIndex].GetValue())

	var l Lang
	errJ = json.Unmarshal(b, &l)
	if errJ != nil {
		log.Fatal(errJ)
	}

	return &l
}
