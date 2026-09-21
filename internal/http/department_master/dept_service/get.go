package deptservice

import (
	"context"
	"fmt"

	"github.com/rajeshbond/smart/internal/auth"
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
	claims *auth.UserClaims,
	id int64,
) (*deptdto.Department, error) {

	if claims == nil {
		return nil, fmt.Errorf(
			"authentication claims are required",
		)
	}

	if id <= 0 {
		return nil, ErrInvalidID
	}

	return s.DeptStore.GetByID(
		ctx,
		id,
	)

}
