package deptservice

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/auth"
	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

func (s *DeptService) CreateDepartmentTx(
	ctx context.Context,
	tx *sql.Tx,
	userID, tenantID int64,
	req *deptdto.CreateDepartmentRequest,
) (*deptdto.Department, error) {

	if tx == nil {
		return nil, fmt.Errorf("transaction is required")
	}

	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	if tenantID <= 0 {
		return nil, fmt.Errorf("invalid tenant ID")
	}

	if req == nil || strings.TrimSpace(req.Department) == "" {
		return nil, fmt.Errorf("department is required")
	}

	if len(strings.TrimSpace(req.Department)) > 100 {
		return nil, fmt.Errorf("department must not exceed 100 characters")
	}

	return s.DeptStore.CreateDepartmentStore(
		ctx,
		tx,
		userID,
		tenantID,
		req,
	)
}

// ============================================================
// CREATE
// ============================================================

// func (s *DeptService) CreateDepartemntService(ctx context.Context, userID, tenantID int64, req *deptdto.CreateDepartmentRequest) (*deptdto.Department, error) {
// 	// 1. Validate authentication
// 	if userID <= 0 {
// 		return nil, ErrUserIDRequired
// 	}

// 	if tenantID <= 0 {
// 		return nil, ErrTenantIDRequired
// 	}

// 	// 2. Validate request
// 	if req == nil {
// 		return nil, ErrDepartmentRequired
// 	}

// 	req.Department = strings.TrimSpace(req.Department)

// 	if req.Department == "" {
// 		return nil, ErrDepartmentRequired
// 	}

// 	if len(req.Department) > 100 {
// 		return nil, fmt.Errorf(
// 			"department must not exceed 100 characters",
// 		)
// 	}

// 	// 3. Begin Transmission
// 	tx, err := s.DeptStore.BeginTx(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("begin department transaction: %w", err)
// 	}

// 	// 4. Ensure rollback if transaction is not committed
// 	committed := false

// 	defer func() {
// 		if !committed {
// 			_ = tx.Rollback()
// 		}
// 	}()

// 	// 5. create department

// 	result, err := s.DeptStore.CreateDepartmentStore(
// 		ctx,
// 		tx,
// 		userID,
// 		tenantID,
// 		req,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	// 6. Commit transaction
// 	if err := tx.Commit(); err != nil {
// 		return nil, fmt.Errorf(
// 			"commit department transaction: %w",
// 			err,
// 		)
// 	}

// 	committed = true

// 	return result, nil

// }

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

	if claims.TenantID <= 0 {
		return nil, ErrTenantIDRequired
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
				claims.TenantID,
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
