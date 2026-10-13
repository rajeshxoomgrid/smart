package partstore

import "database/sql"

type PartStore struct {
	db *sql.DB
}

func NewPartStore(db *sql.DB) *PartStore {
	return &PartStore{db: db}
}
