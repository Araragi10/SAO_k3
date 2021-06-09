package wotoSudo

import (
	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToSudoList(m []primitive.M) interfaces.SudoList {
	if m == nil {
		return nil
	}
	sudoList := make(sudoList, wv.BaseIndex)
	var sudo sudoInfo

	for _, current := range m {
		sudo = sudoInfo{
			M: current,
		}

		sudoList = append(sudoList, sudo)
	}

	return &sudoList
}
