package wotoGP

type gpError uint8

const (
	usernameWrong gpError = 1
)

const (
	packErrSign         = "wotoGP package: "
	usernameEmptyString = packErrSign +
		"username cannot be empty"
	settingsNilString = packErrSign +
		"helpers.go: settings is nil"
	gClientNilString = packErrSign +
		"gClient should not be empty"
	resolvedPeerNilString = packErrSign +
		"the resolved peer is nil"
	userChZeroString = packErrSign +
		"found zero users and chats. check the username again"
	userZeroString = packErrSign +
		"found zero users and chats. check the username again"
	convertFailedString = packErrSign +
		"failed to convert recievied data"
)
