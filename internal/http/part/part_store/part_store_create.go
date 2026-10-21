package partstore

import (
	"context"
	"database/sql"

	partdto "github.com/rajeshbond/smart/internal/http/part/part_dto"
)

func (s *PartStore) Create(
	ctx context.Context,
	tx *sql.Tx,
	tenantID int64,
	req partdto.CreatePartRequest,
) (*partdto.PartResponse, error) {

	isActive := true

	// Use the supplied value if the client provided one.
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	query := `
		INSERT INTO part (
			tenant_id,
			part_no,
			product_name,
			color,
			rm_grade,
			part_wt,
			description,
			is_active
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING
			id, tenant_id, part_no, product_name,
			color, rm_grade, part_wt, description,
			is_active, is_deleted, deleted_by,
			created_at, updated_at
	`

	part, err := scanPart(
		tx.QueryRowContext(
			ctx,
			query,
			tenantID,
			req.PartNo,
			req.ProductName,
			req.Color,
			req.RMGrade,
			req.PartWt,
			req.Description,
			isActive,
		),
	)
	if err != nil {
		return nil, translateDBError(err)
	}

	// scanPart returns a *dto.PartResponse.
	return part, nil
}
