// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"log"
	"strings"

	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/pTools"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CmdHandler is the c
type CmdHandler func(*tg.Message, *pTools.EventArgs)

// IsCommand will check if the _text is a command or not.
func IsCommand(text *string) bool {
	return isCommand1(text) || isCommand2(text)
}

// isCommand1 will check if the _text
// has first prefex or not.
func isCommand1(text *string) bool {
	return strings.HasPrefix(*text, wv.COMMAND_PREFIX1)
}

// isCommand2 will check if the _text
// has first prefex or not.
func isCommand2(text *string) bool {
	return strings.HasPrefix(*text, wv.COMMAND_PREFIX2)
}

func HandleCommand(message *tg.Message) {
	cmdListInit()
	eventArg, err := pTools.ParseArg(message.Text)
	if err != nil {
		log.Println("an error in parsing the args:", err)
		return
	}

	if eventArg == nil {
		log.Println("unexpected error, event arg is nil")
		return
	}

	event, ok := cmdList[eventArg.GetCommand()]
	if event != nil && ok {
		event(message, eventArg)
	}
}

//func main() {
//    input := `xxxxx:yyyyy:zzz.aaa.bbb.cc:dd:ee:ff`
//    a := strings.FieldsFunc(input, Split)
//    t := Target{a[0], a[1], a[2], a[3], a[4], a[5], a[6]}
//    fmt.Println(t) // {xxxxx yyyyy zzz aaa bbb cc dd}
//}
//func Split(r rune) bool {
//    return r == ':' || r == '.'
//}
