package permissionservice

import (
	"context"
	"errors"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

// ============================================================
// CREATE
// ============================================================

func (s *PermissionService) Create(
	ctx context.Context,
	req *permissiondto.CreatePermissionRequest,
) (*permissiondto.Permission, error) {

	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if req.UserID <= 0 {
		return nil, errors.New("user_id must be greater than zero")
	}

	if req.RoleID <= 0 {
		return nil, errors.New("role_id must be greater than zero")
	}

	if req.DepartmentID <= 0 {
		return nil, errors.New("department_id must be greater than zero")
	}

	return s.PermissionStore.Create(
		ctx,
		req,
	)
}
