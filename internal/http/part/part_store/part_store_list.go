package partstore

import (
	"context"
	"fmt"
	"strings"

	partdto "github.com/rajeshbond/smart/internal/http/part/part_dto"
)

func (s *PartStore) GetPartList(
	ctx context.Context,
	tenantID int64,
	req *partdto.ListPartsRequest,
) (*partdto.ListPartsResponse, error) {

	where := []string{
		"tenant = $1",
		"is_deleted = FALSE",
	}
	args := []any{tenantID}

	// Search by part number or product name
	if req.Search != "" {
		args = append(args, "%"+req.Search+"%")
		n := len(args)

		where = append(where,
			fmt.Sprintf("(part_no ILIKE $%d OR product_name ILIKE $%d)", n, n),
		)
	}

	// Filter by active status
	if req.IsActive != nil {
		args = append(args, *req.IsActive)
		where = append(where,
			fmt.Sprintf("is_active = $%d", len(args)),
		)
	}

	whereSQL := strings.Join(where, " AND ")

	// Get total matching records
	var total int64

	countQuery := `
		SELECT COUNT(*)
		FROM part
		WHERE ` + whereSQL

	if err := s.db.QueryRowContext(
		ctx,
		countQuery,
		args...,
	).Scan(&total); err != nil {
		return nil, fmt.Errorf("count parts: %w", err)
	}

	// Calculate pagination
	offset := (req.Page - 1) * req.PageSize

	// Add pagination arguments
	listArgs := append([]any{}, args...)
	listArgs = append(listArgs, req.PageSize, offset)

	limitParam := len(args) + 1
	offsetParam := len(args) + 2

	// Fetch parts
	query := fmt.Sprintf(`
		SELECT
			*
		FROM part
		WHERE %s
		ORDER BY part_no ASC
		LIMIT $%d OFFSET $%d
	`, whereSQL, limitParam, offsetParam)

	rows, err := s.db.QueryContext(ctx, query, listArgs...)
	if err != nil {
		return nil, fmt.Errorf("list parts: %w", err)
	}
	defer rows.Close()

	items := make([]partdto.PartResponse, 0)

	for rows.Next() {
		part, err := scanPart(rows)
		if err != nil {
			return nil, fmt.Errorf("scan part: %w", err)
		}
		items = append(items, *part)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate parts: %w", err)
	}

	// Calculate total pages
	var totalPages int64
	if total > 0 {
		totalPages = (total + int64(req.PageSize) - 1) /
			int64(req.PageSize)
	}

	return &partdto.ListPartsResponse{
		Items:      items,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}
