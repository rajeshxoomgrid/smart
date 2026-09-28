package permissionservice

import (
	"context"

	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

type DeptProvider interface {
	GetByID(ctx context.Context, id int64) (*deptdto.Department, error)
	etByName(ctx context.Context, name string) (*deptdto.Department, error)
}
