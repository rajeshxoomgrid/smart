package partservice

import (
	"context"
	"fmt"

	partdto "github.com/rajeshbond/smart/internal/http/part/part_dto"
	partserviceconstant "github.com/rajeshbond/smart/internal/http/part/part_service/partServiceConstant"
)

func (ser *PartService) CreatePartService(
	ctx context.Context,
	req partdto.CreatePartRequest,
	tenantID int64,
) (*partdto.PartResponse, error) {

	// 1. Validate tenant ID.
	if tenantID <= 0 {
		return nil, fmt.Errorf(
			"%w: tenant ID is required",
			partserviceconstant.ErrInvalidRequest,
		)
	}

	// 2. Normalize optional strings.
	nornalizeNullableStrings(
		req.Color,
		req.RMGrade,
		req.Description,
	)

	// 3. Validate request fields.
	if err := valodateField(
		req.PartNo,
		req.ProductName,
		req.Color,
		req.RMGrade,
		req.PartWt,
	); err != nil {
		return nil, err
	}

	// 4. Begin database transaction.
	tx, err := ser.PartStore.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	// Roll back if the transaction has not been committed.
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	// 5. Insert the part using the transaction.
	// Replace CreatePart with your actual PartStore insert method.
	result, err := ser.PartStore.Create(ctx, tx, tenantID, req)
	if err != nil {
		return nil, fmt.Errorf("create part: %w", err)
	}

	// 6. Commit the transaction.
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	committed = true

	// 7. Return the created part.
	return result, nil
}

// Normalize optional strings.
