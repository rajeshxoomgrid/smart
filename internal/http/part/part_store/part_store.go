package partstore

import (
	"context"
	"database/sql"
)

type PartStore struct {
	db *sql.DB
}

func NewPartStore(db *sql.DB) *PartStore {
	return &PartStore{db: db}
}

func (s *PartStore) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, nil)
}
