// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoGP

// the flags for the wotoMorse plugin.
// please notice that all flags should begin with
// prefex --, but in the pTools, this prefex will be
// removed.
const (
	PV_FLAG      = "pv"
	PRIVATE_FLAG = "private"
	DEL_FLAG     = "del"    // won't work if message is not replied
	DELETE_FLAG  = "delete" // won't work if message is not replied
	USER_FLAG    = "user"   // won't work if message is not replied
	USR_FLAG     = "usr"    // won't work if message is not replied
)
