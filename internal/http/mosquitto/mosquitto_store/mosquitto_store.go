package mosquittostore

import "database/sql"

type MosquittoStore struct {
	db *sql.DB
}

func NewMosquittoStore(db *sql.DB) *MosquittoStore {
	return &MosquittoStore{db: db}
}
