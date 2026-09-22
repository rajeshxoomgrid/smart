package users

import (
	"context"

	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

type DeptProvider interface {
	GetByID(ctx context.Context, id int64) (*deptdto.Department, error)
}
