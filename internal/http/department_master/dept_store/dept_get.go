package deptstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

func (s *DeptStore) GetByID(ctx context.Context, id int64) (*deptdto.Department, error) {

	var result deptdto.Department

	err := s.db.QueryRowContext(
		ctx,
		queryGetDepartmentByID,
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
			return nil, errors.New("department not found")
		}

		return nil, fmt.Errorf(
			"get department by id %d: %w",
			id,
			err,
		)
	}

	return &result, nil
}

func (s *DeptStore) GetByName(
	ctx context.Context,
	name string,
) (*deptdto.Department, error) {

	var result deptdto.Department

	err := s.db.QueryRowContext(
		ctx,
		queryGetDepartmentByName,
		name,
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
			return nil, errors.New("department not found")
		}

		return nil, fmt.Errorf(
			"get department by name %s: %w",
			name,
			err,
		)
	}

	return &result, nil
}
