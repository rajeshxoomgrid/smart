package moldchangeservice

import (
	moldchangestore "github.com/rajeshbond/smart/internal/http/mold_change_register/mold_change_store"
)

type MoldChangeService struct {
	MoldChangeStore *moldchangestore.MoldChangeStore
}

func NewMoldChangeService(moldChangeStoe *moldchangestore.MoldChangeStore) *MoldChangeService {
	return &MoldChangeService{MoldChangeStore: moldChangeStoe}
}
