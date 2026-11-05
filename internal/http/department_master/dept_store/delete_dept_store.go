package deptstore

import (
	"context"
	"database/sql"
	"fmt"
)

// ============================================================
// SOFT DELETE
// ============================================================

func (s *DeptStore) SoftDelete(
	ctx context.Context,
	tx *sql.Tx,
	id int64,
	userID int64,
) error {

	result, err := tx.ExecContext(
		ctx,
		querySoftDeleteDepartment,
		userID,
		id,
	)

	if err != nil {
		return fmt.Errorf(
			"delete department: %w",
			err,
		)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf(
			"delete department rows affected: %w",
			err,
		)
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
