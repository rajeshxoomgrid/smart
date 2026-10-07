package immhttpstore

import (
	"context"
	"database/sql"
)

// ============================================================
// TRANSACTION
// ============================================================

// ============================================================
// TRANSACTION
// ============================================================

func (s *ImmHttpStore) WithTransaction(
	ctx context.Context,
	fn func(tx *sql.Tx) error,
) error {

	tx, err := s.db.BeginTx(
		ctx,
		&sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		},
	)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}

// func (s *ImmHttpStore) WithTransaction(
// 	ctx context.Context,
// 	fn func(tx *sql.Tx) error,
// ) error {

// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	if err := fn(tx); err != nil {
// 		_ = tx.Rollback()
// 		return err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		_ = tx.Rollback()
// 		return err
// 	}

// 	return nil
// }
