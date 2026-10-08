package immstore

import "database/sql"

type ImmStore struct {
	db *sql.DB
}

func NewImmStore(db *sql.DB) *ImmStore {
	return &ImmStore{db: db}
}
