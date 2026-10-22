package partstore

import (
	"context"
	"database/sql"
	"fmt"

	partconstant "github.com/rajeshbond/smart/internal/http/part/part_constant"
)

func (s *PartStore) PartSoftDelete(ctx context.Context, tx *sql.Tx, tenantID, id, deleted_by int64) error {
	softDelereQuery := `
		UPDATE part 
		SET 
			is_deleted = TRUE
			updated_at = NOW()
		WHERE tenant_id = $1
		AND id = $2
		AND is_deleted = FALSE
	`

	result, err := tx.ExecContext(ctx, softDelereQuery, tenantID, id, deleted_by)
	if err != nil {
		return fmt.Errorf("check deleted part:%w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted parts:%w", err)
	}
	if rows == 0 {
		return partconstant.ErrNotFound
	}

	return nil
}
