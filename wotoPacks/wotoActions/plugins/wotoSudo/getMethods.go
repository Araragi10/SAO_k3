package wotoSudo

import (
	"time"

	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
type SudoInfo interface {
	GetID() int64
	GetNickname() string
	GetDate() time.Time
}

type SudoList interface {
	GetSudo(int64) *SudoInfo
	GetListP() []primitive.M
}



*/
//---------------------------------------------------------

func (s *sudoInfo) GetID() int64 {
	id, ok := s.M[wv.SudoNick].(int64)
	if !ok {
		return wv.BaseIndex
	}

	return id
}

func (s *sudoInfo) GetNickname() string {
	n, ok := s.M[wv.SudoNick].(string)
	if !ok {
		return wv.EMPTY
	}

	return n
}

func (s *sudoInfo) GetDate() time.Time {
	t, ok := s.M[wv.SudoDate].(time.Time)
	if !ok {
		return time.Time{}
	}

	return t
}

func (s *sudoInfo) GetAsP() *primitive.M {
	return &s.M
}

//---------------------------------------------------------

func (l *sudoList) GetSudo(id *int64) interfaces.SudoInfo {
	for _, s := range *l {
		if s.GetID() == *id {
			return &s
		}
	}

	return nil
}

func (l *sudoList) Contains(id *int64) bool {
	for _, s := range *l {
		if s.GetID() == *id {
			return true
		}
	}

	return false
}

func (l *sudoList) GetListP() []primitive.M {
	final := make([]primitive.M, wv.BaseIndex)
	for _, s := range *l {
		final = append(final, *s.GetAsP())
	}

	return final
}

//---------------------------------------------------------
//---------------------------------------------------------
