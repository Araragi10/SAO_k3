package wotoSudo

type sudoErr uint16

const (
	AnotInList     sudoErr = 1
	AinvalidId     sudoErr = 2
	AalreadyInList sudoErr = 3
)
