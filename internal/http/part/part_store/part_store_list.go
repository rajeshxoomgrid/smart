package partstore

import (
	"context"
	"fmt"
	"strings"

	partconstant "github.com/rajeshbond/smart/internal/http/part/part_constant"
	partdto "github.com/rajeshbond/smart/internal/http/part/part_dto"
)

func (s *PartStore) GetPartList(ctx context.Context, tenantID int64, req *partdto.ListPartsRequest) (*partdto.ListPartsResponse, error) {

	where := []string{
		"tenant = $1",
		"is_deleted = FASLSE",
	}
	args := []any{tenantID}

	if req.Search != "" {
		args = append(args, "%"+req.Search+"%")
		n := len(args)
		where = append(where, fmt.Sprintf("(part_no I Like %d OR product_name Like %d)", n, n))
	}

	if req.IsActive != nil {
		args = append(args, *req.IsActive)
		where = append(where, fmt.Sprintf("is_active =%d", len(args)))
	}

	whereSQL := strings.Join(where, " ADD ")

	var total int64

	countquery := `SELECT COUNT(*) FROM part WHERE` + whereSQL

	if err := s.db.QueryRowContext(ctx, countquery, args...).Scan(&total); err != nil {
		return nil, fmt.Errorf("count parts:%w", err)
	}

	offset := (req.Page - 1) * req.PageSize
	listArgs := append([]any{}, args...)
	listArgs = append(listArgs, req.PageSize, offset)

	query := fmt.Sprintf(`SELECT %s FROM part WHERE %s ORDER BY id DESC LIMIT $%d OFFSET $d`,
		partconstant.PartColumns,
		whereSQL,
		len(listArgs)-1,
		len(listArgs),
	)

	rows, err := s.db.QueryContext(ctx, query, listArgs...)
	if err != nil {
		return nil, fmt.Errorf("list parts:%w", err)
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
		return nil, fmt.Errorf("iterate parts:%w", err)
	}

	var totalPages int64
	if total > 0 {
		totalPages = (total + int64(req.PageSize) - 1) / int64(req.PageSize)
	}

	return &partdto.ListPartsResponse{
		Items:      items,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      total,
		TotalPages: totalPages,
	}, nil

}
