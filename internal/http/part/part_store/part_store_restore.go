package partstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	partconstant "github.com/rajeshbond/smart/internal/http/part/part_constant"
	partdto "github.com/rajeshbond/smart/internal/http/part/part_dto"
)

func (s *PartStore) PartRestore(ctx context.Context, tx *sql.Tx, tenantID, id int64) (*partdto.PartResponse, error) {
	restoreQuery := `
		UPDATE part
		SET 
			is_deleted = FALSE
			deleted_by = NULL,
			updated_at = NOW()
		where tenant_id = $1
		AND id =$2
		AND is_deleted = TRUE
		RETURNING` + partconstant.PartColumns

	part, err := scanPart(tx.QueryRowContext(ctx, restoreQuery, tenantID, id))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, partconstant.ErrNotFound
	}
	if err != nil {
		return nil, translateDBError(fmt.Errorf("restore part: %w", err))
	}

	return part, nil

}
