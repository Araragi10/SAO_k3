// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoSudo

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Araragi10/SAO_k3/wotoPacks/appSettings"
	wa "github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/common"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoActions/plugins/pTools"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoMD"
	ws "github.com/Araragi10/SAO_k3/wotoPacks/wotoSecurity/wotoStrings"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoSys"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoValues/tgMessages"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoValues/tgMessages/tgConst"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Sudo_handler(message *tg.Message, args *pTools.EventArgs) {
	// set the first element of args to empty,
	// because pTools.Arg also contains the command itself,
	// however we don't want that here.
	//args[wv.BaseIndex] = wv.EMPTY
	isMain := appSettings.GetExisting().IsMainSudo(message.From.ID)

	if args.HasFlag(SYS_FLAG) {
		s, err := wotoSys.GetSysInfo()
		str := ""
		if err != nil {
			str = err.Error()
		} else {
			data, err := json.MarshalIndent(s, "", "  ")
			if err != nil {
				str = err.Error()
			} else {
				str = string(data)
			}
		}
		appSettings.GetExisting().SendSudo(str)
		return
	}

	if args.HasFlag(ADD_FLAG) && isMain {
		addSudo(message, args)
	} else if args.HasFlag(REMOVE_FLAG, REM_FLAG, RM_FLAG) && isMain {
		remSudo(message, args)
	} else if args.HasFlag(LS_FLAG, LIST_FLAG) {
		listSudo(message, args)
	}
}

func listSudo(message *tg.Message, args *pTools.EventArgs) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

}

func addSudo(message *tg.Message, args *pTools.EventArgs) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	var full string
	var id int64
	var nick string
	var err error
	is_reply := message.ReplyToMessage != nil
	send_pv := args.HasFlag(PV_FLAG, PRIVATE_FLAG)

	if is_reply {
		if message.ReplyToMessage.From != nil {
			id = message.ReplyToMessage.From.ID
			nick = message.ReplyToMessage.From.FirstName
		} else {
			full = args.GetAsStringOrRaw(USER_FLAG, USR_FLAG)
			id, err = strconv.ParseInt(full, wv.BaseTen, wv.Base64Bit)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		full = args.GetAsStringOrRaw(USER_FLAG, USR_FLAG)
		// don't foget to trim the spaces.
		full = strings.Trim(full, wv.SPACE_VALUE)
		id, err = strconv.ParseInt(full, wv.BaseTen, wv.Base64Bit)
		if err != nil {
			log.Println(err)
			invalid_id(message, full, send_pv)
			return
		}
		if id == wv.BaseIndex {
			log.Println(ID_CANNOT_BE_ZERO)
		}
	}

	res := settings.AddSudo(id, nick)
	switch res {
	case wa.SUCCESS:
		added_notice(message, id, send_pv)
	case wa.CANCELED:
		already_in_list(message, id, send_pv)
	case wa.FAILED:
		log.Println(ADD_SUDO_ERR)
	}
}

func remSudo(message *tg.Message, args *pTools.EventArgs) {
	settings := appSettings.GetExisting()
	var full string
	var id int64
	var err error
	is_reply := message.ReplyToMessage != nil
	send_pv := args.HasFlag(PV_FLAG, PRIVATE_FLAG)

	if is_reply {
		if message.ReplyToMessage.From != nil {
			id = message.ReplyToMessage.From.ID
		} else {
			invalid_id(message, args.GetAsStringOrRaw(USER_FLAG, USR_FLAG), send_pv)
			return
		}
	} else {
		full = args.GetAsStringOrRaw(USER_FLAG, USR_FLAG)
		// don't foget to trim the spaces.
		full = strings.Trim(full, wv.SPACE_VALUE)
		id, err = strconv.ParseInt(full, wv.BaseTen, wv.Base64Bit)
		if err != nil {
			log.Println(err)
			invalid_id(message, full, send_pv)
			return
		}
		if id == wv.BaseIndex {
			log.Println(ID_CANNOT_BE_ZERO)
		}
	}

	res := settings.RemSudo(id)
	switch res {
	case wa.SUCCESS:
		removed_notice(message, id, send_pv)
	case wa.CANCELED:
		not_in_list(message, id, send_pv)
	case wa.FAILED:
		log.Println(REM_SUDO_ERR)
	}
}

func added_notice(message *tg.Message, id int64, pv bool) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	reply := message.ReplyToMessage != nil

	api := settings.GetAPI()
	if api == nil {
		return
	}

	var msg tg.MessageConfig
	var text string

	if reply {
		text = message.ReplyToMessage.From.FirstName
	} else {
		text = strconv.FormatInt(id, wv.BaseTen)
	}

	if pv {
		str := ADDED_TO_SUDO_LIST_MPV
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.From.ID, str)
	} else {
		str := ADDED_TO_SUDO_LIST_M
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.Chat.ID, str)
	}

	// for fixing: Bad Request: replied message not found
	if !pv {
		msg.ReplyToMessageID = message.MessageID
	}

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func removed_notice(message *tg.Message, id int64, pv bool) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	reply := message.ReplyToMessage != nil &&
		!ws.IsEmpty(&message.ReplyToMessage.Text)

	api := settings.GetAPI()
	if api == nil {
		return
	}

	var msg tg.MessageConfig
	var text string

	if reply {
		text = message.ReplyToMessage.From.FirstName
	} else {
		text = strconv.FormatInt(id, wv.BaseTen)
	}

	if pv {
		str := REMOVED_SUDO_MPV
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.From.ID, str)
	} else {
		str := REMOVED_SUDO_M
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.Chat.ID, str)
	}

	// for fixing: Bad Request: replied message not found
	if !pv {
		msg.ReplyToMessageID = message.MessageID
	}

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func not_in_list(message *tg.Message, id int64, pv bool) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	reply := message.ReplyToMessage != nil

	api := settings.GetAPI()
	if api == nil {
		return
	}

	var msg tg.MessageConfig
	var text string

	if reply {
		text = message.ReplyToMessage.From.FirstName
	} else {
		text = strconv.FormatInt(id, wv.BaseTen)
	}

	if pv {
		str := NOT_IN_SUDO_LIST_MPV
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.From.ID, str)
	} else {
		str := NOT_IN_SUDO_LIST_M
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.Chat.ID, str)
	}

	// for fixing: Bad Request: replied message not found
	if !pv {
		msg.ReplyToMessageID = message.MessageID
	}

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func already_in_list(message *tg.Message, id int64, pv bool) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	reply := message.ReplyToMessage != nil &&
		!ws.IsEmpty(&message.ReplyToMessage.Text)

	api := settings.GetAPI()
	if api == nil {
		return
	}

	var msg tg.MessageConfig
	var text string

	if reply {
		text = message.ReplyToMessage.From.FirstName
	} else {
		text = strconv.FormatInt(id, wv.BaseTen)
	}

	if pv {
		str := ALREADY_SUDO_MPV
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.From.ID, str)
	} else {
		str := ALREADY_SUDO_M
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.Chat.ID, str)
	}

	// for fixing: Bad Request: replied message not found
	if !pv {
		msg.ReplyToMessageID = message.MessageID
	}

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func invalid_id(message *tg.Message, id string, pv bool) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	var msg tg.MessageConfig

	if pv {
		str := fmt.Sprintf(INVALID_ID_SUDO_MPV, id)
		str = wotoMD.GetNormal(str).ToString()
		msg = tg.NewMessage(message.From.ID, str)
	} else {
		str := fmt.Sprintf(INVALID_ID_SUDO_M, id)
		str = wotoMD.GetNormal(str).ToString()
		msg = tg.NewMessage(message.Chat.ID, str)
	}

	// for fixing: Bad Request: replied message not found
	if !pv {
		msg.ReplyToMessageID = message.MessageID
	}

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}
