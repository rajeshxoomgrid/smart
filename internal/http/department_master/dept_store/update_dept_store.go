package deptstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

// ============================================================
// UPDATE
// ============================================================

func (s *DeptStore) Update(
	ctx context.Context,
	tx *sql.Tx,
	id int64,
	userID int64,
	req *deptdto.UpdateRequest,
) (*deptdto.Department, error) {

	department := strings.TrimSpace(
		req.Department,
	)

	var result deptdto.Department

	err := tx.QueryRowContext(
		ctx,
		queryUpdateDepartment,
		department,
		userID,
		id,
	).Scan(
		&result.ID,
		&result.Department,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.IsDeleted,
		&result.DeletedAt,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		var pqErr *pq.Error

		if errors.As(err, &pqErr) &&
			pqErr.Code == "23505" {

			return nil, ErrDuplicate
		}

		return nil, fmt.Errorf(
			"update department: %w",
			err,
		)
	}

	return &result, nil
}
