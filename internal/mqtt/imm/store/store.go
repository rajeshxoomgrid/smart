package store

import (
	"database/sql"
)

// ============================================================
// IMM PRODUCTION STORE
// ============================================================

type ImmProductionStore struct {
	db *sql.DB
}

// ============================================================
// NEW IMM PRODUCTION STORE
// ============================================================

func NewImmProductionStore(
	db *sql.DB,
) *ImmProductionStore {

	// if db == nil {
	// 	panic("IMM production store database connection cannot be nil")
	// }

	return &ImmProductionStore{
		db: db,
	}
}
