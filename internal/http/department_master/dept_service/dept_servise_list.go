package deptservice

import (
	"context"
	"fmt"

	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

func (ser *DeptService) ListAllDepartemntTenantService(ctx context.Context, tenantID int64) ([]deptdto.Department, error) {
	if tenantID <= 0 {
		return nil, fmt.Errorf("invalid tenant ID")
	}

	departments, err := ser.DeptStore.ListTenantAllDepartmentStore(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("list departemts:%w", err)
	}

	return departments, nil
}
