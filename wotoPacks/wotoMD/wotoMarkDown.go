// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMD

import (
	"strconv"

	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

type wotoMarkDown struct {
	_value string
}

func GetNormal(value string) interfaces.WMarkDown {
	final := repairValue(value)
	return toWotoMD(final)
}

func GetBold(value string) interfaces.WMarkDown {
	final := repairValue(value)
	final = string(CHAR_S3) + final + string(CHAR_S3)
	return toWotoMD(final)
}

func GetItalic(value string) interfaces.WMarkDown {
	final := repairValue(value)
	final = string(CHAR_S4) + final + string(CHAR_S4)
	return toWotoMD(final)
}

func GetMono(value string) interfaces.WMarkDown {
	final := repairValue(value)
	final = string(CHAR_S16) + final + string(CHAR_S16)
	return toWotoMD(final)
}

func GetHyperLink(text string, url string) interfaces.WMarkDown {
	fText := repairValue(text)
	fUrl := repairValue(url)
	final := string(CHAR_S7) + fText + string(CHAR_S8) +
		string(CHAR_S9) + fUrl + string(CHAR_S10)
	return toWotoMD(final)
}

// GetUserMention will give you a mentioning style username with the
// specified text.
// WARNING: you don't need to repair text before sending it as first arg,
// this function will check it itself.
func GetUserMention(text string, userID int64) interfaces.WMarkDown {
	final := repairValue(text)
	final = string(CHAR_S7) + final +
		string(CHAR_S8) +
		string(CHAR_S9) + TG_USER_ID +
		strconv.FormatInt(userID, wv.BaseTen) +
		string(CHAR_S10)

	return toWotoMD(final)
}

func GetUserMentionByUsername(text, username string) interfaces.WMarkDown {
	final := repairValue(text)
	final = string(CHAR_S7) + final +
		string(CHAR_S8) +
		string(CHAR_S9) + TG_USER_UN +
		username +
		string(CHAR_S10)

	return toWotoMD(final)
}

func IsSpecial(r rune) bool {
	for _, current := range _sChars {
		if r == current {
			return true
		}
	}
	return false
}

func IsEnglish(r rune) bool {
	if r >= LetterA && r <= LetterZ {
		return true
	} else {
		return r >= CapLetterA && r <= CapLetterZ
	}
}

func IsNum(r rune) bool {
	return r >= LetterZero || r <= LetterNine
}

func IsNumOrEng(r rune) bool {
	return IsNum(r) || IsEnglish(r)
}

func toWotoMD(value string) interfaces.WMarkDown {
	if wotoStrings.IsEmpty(&value) {
		return nil
	}

	myMD := wotoMarkDown{
		_value: value,
	}
	return &myMD
}

func repairValue(value string) string {
	finally := wv.EMPTY
	//escape := false
	//lasEscape := false
	//escapeCount := wv.BaseIndex
	for _, current := range value {
		//escape = (current == CHAR_S1)
		if IsSpecial(current) {
			finally += string(CHAR_S1)
		}
		finally += string(current)
	}
	return finally
}
