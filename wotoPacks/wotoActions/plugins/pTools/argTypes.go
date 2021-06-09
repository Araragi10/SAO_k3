// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package pTools

type FlagType uint8

const (
	NoneFlagType   FlagType = iota
	BoolFlagType   FlagType = iota + 1
	StringFlagType FlagType = iota + 2
	UInt8FlagType  FlagType = iota + 3
	UInt16FlagType FlagType = iota + 4
	UInt32FlagType FlagType = iota + 5
	UInt64FlagType FlagType = iota + 6
	Int8FlagType   FlagType = iota + 7
	Int16FlagType  FlagType = iota + 8
	Int32FlagType  FlagType = iota + 9
	Int64FlagType  FlagType = iota + 10
)

const (
	NoneTypeStr   = "None"
	BoolTypeStr   = "bool"
	StringTypeStr = "string"
	UInt8TypeStr  = "uint8"
	UInt16TypeStr = "uint16"
	UInt32TypeStr = "uint32"
	UInt64TypeStr = "uint64"
	Int8TypeStr   = "int8"
	Int16TypeStr  = "int16"
	Int32TypeStr  = "int32"
	Int64TypeStr  = "int64"
)
