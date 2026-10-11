package service

import (
	"github.com/rajeshbond/smart/internal/mqtt/imm/store"
)

// ============================================================
// IMM SERVICE
// ============================================================

type ImmService struct {
	ImmProductionStore *store.ImmProductionStore
}

func NewImmService(
	immStore *store.ImmProductionStore,
) ImmService {
	return ImmService{
		ImmProductionStore: immStore,
	}
}
