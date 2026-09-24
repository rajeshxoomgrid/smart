package deptservice

import (
	"context"

	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

// import (
// 	"context"
// 	"fmt"

// 	"github.com/rajeshbond/smart/internal/auth"
// 	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
// )

// // ============================================================
// // GET
// // ============================================================

func (s *DeptService) GetByID(
	ctx context.Context,
	id int64,
) (*deptdto.Department, error) {

	if id <= 0 {
		return nil, ErrInvalidID
	}

	return s.DeptStore.GetByID(
		ctx,
		id,
	)
}

func (s *DeptService) GetByName(
	ctx context.Context,
	name string,
) (*deptdto.Department, error) {

	if name == "" {
		return nil, ErrInvalidName
	}

	return s.DeptStore.GetByName(
		ctx,
		name,
	)
}
