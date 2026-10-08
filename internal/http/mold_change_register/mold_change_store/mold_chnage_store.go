package moldchangestore

import "database/sql"

type MoldChangeStore struct {
	db *sql.DB
}

func NewMoldChangeStore(db *sql.DB) *MoldChangeStore {
	return &MoldChangeStore{db: db}
}
