package wotoSudo

import "go.mongodb.org/mongo-driver/bson/primitive"

type sudoInfo struct {
	primitive.M
}

type sudoList []sudoInfo
