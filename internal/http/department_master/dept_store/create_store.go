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
// CREATE
// ============================================================

func (s *DeptStore) Create(
	ctx context.Context,
	tx *sql.Tx,
	userID int64,
	req *deptdto.CreateDepartmentRequest,
) (*deptdto.Department, error) {

	department := strings.ToLower(
		strings.TrimSpace(req.Department),
	)

	var result deptdto.Department

	err := tx.QueryRowContext(
		ctx,
		queryCreateDepartment,
		department,
		userID,
		userID,
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

		if isDuplicateError(err) {
			return nil, ErrDuplicate
		}

		return nil, fmt.Errorf(
			"create department: %w",
			err,
		)
	}

	return &result, nil
}

// ============================================================
// DUPLICATE ERROR
// ============================================================

func isDuplicateError(err error) bool {

	var pqErr *pq.Error

	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505"
	}

	return false
}
