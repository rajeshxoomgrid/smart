package immhttpstore

import "database/sql"

type ImmHttpStore struct {
	db *sql.DB
}

func NewImmHttpStore(db *sql.DB) *ImmHttpStore {
	return &ImmHttpStore{db: db}
}
