package partstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	partconstant "github.com/rajeshbond/smart/internal/http/part/part_constant"
	partdto "github.com/rajeshbond/smart/internal/http/part/part_dto"
)

func (s *PartStore) PartUpdate(ctx context.Context, tx *sql.Tx, tenantID, id int64, req partdto.UpdatePartRequest) (*partdto.PartResponse, error) {
	updareQuery := `
		UPDATE part
		SET
			part_no = $3
			part_name=$4,
			color = $5,
			rm_grade = $6,
			part_wt = $7,
			description =$8,
			is_active = $9,
			updated_at = NOW()
		WHERE tenant_id = $1
			AND id = $2 
			AND is_deleted = FALSE
			RETURNING ` + partconstant.PartColumns

	part, err := scanPart((tx.QueryRowContext(
		ctx,
		updareQuery,
		tenantID,
		id,
		req.PartNo,
		req.ProductName,
		req.Color,
		req.RMGrade,
		req.PartWt,
		req.Description,
		req.IsActive,
	)))

	if errors.Is(err, sql.ErrNoRows) {
		return nil, partconstant.ErrNotFound
	}

	if err != nil {
		return nil, translateDBError(fmt.Errorf("upate part: %w", err))
	}

	return part, nil
}
