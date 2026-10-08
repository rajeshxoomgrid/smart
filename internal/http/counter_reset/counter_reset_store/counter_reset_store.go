package counterresetstore

import "database/sql"

type CounterResetStore struct {
	db *sql.DB
}

func NewCounterRestStore(db *sql.DB) *CounterResetStore {
	return &CounterResetStore{db: db}
}
