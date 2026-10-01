package permissionstore

import (
	"context"
	"database/sql"
)

// ============================================================
// DELETE - SOFT DELETE
// ============================================================

func (s *PermissionStore) Delete(
	ctx context.Context,
	id int64,
	deletedBy *int64,
) error {

	query := `
		UPDATE public.permission
		SET
			is_deleted = true,
			deleted_by = $2,
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
		  AND is_deleted = false
	`

	result, err := s.db.ExecContext(
		ctx,
		query,
		id,
		deletedBy,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
