package deptservice

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/auth"
	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

// ============================================================
// CREATE
// ============================================================

func (s *DeptService) Create(
	ctx context.Context,
	claims *auth.UserClaims,
	req *deptdto.CreateDepartmentRequest,
) (*deptdto.Department, error) {

	if claims == nil {
		return nil, fmt.Errorf(
			"authentication claims are required",
		)
	}

	if claims.UserID <= 0 {
		return nil, ErrUserIDRequired
	}

	if req == nil {
		return nil, ErrDepartmentRequired
	}

	req.Department = strings.TrimSpace(
		req.Department,
	)

	if req.Department == "" {
		return nil, ErrDepartmentRequired
	}

	if len(req.Department) > 100 {
		return nil, fmt.Errorf(
			"department must not exceed 100 characters",
		)
	}

	var result *deptdto.Department

	err := s.withTransaction(
		ctx,
		func(tx *sql.Tx) error {

			var err error

			result, err = s.DeptStore.Create(
				ctx,
				tx,
				claims.UserID,
				req,
			)

			return err
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
