package immhttpstore

import (
	"context"
	"database/sql"
)

// ============================================================
// TRANSACTION
// ============================================================

func (s *ImmHttpStore) WithTransaction(
	ctx context.Context,
	fn func(tx *sql.Tx) error,
) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
