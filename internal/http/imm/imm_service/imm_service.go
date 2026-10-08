package immservice

import immstore "github.com/rajeshbond/smart/internal/http/imm/imm_store"

type Immservice struct {
	ImmStore *immstore.ImmStore
}

func NewImmStoreImmSerive(ImmStore *immstore.ImmStore) *Immservice {
	return &Immservice{ImmStore: ImmStore}
}
