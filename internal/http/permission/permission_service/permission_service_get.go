package permissionservice

import (
	"context"
	"errors"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

// ============================================================
// GET BY ID
// ============================================================

func (s *PermissionService) GetByID(
	ctx context.Context,
	id int64,
) (*permissiondto.Permission, error) {

	if id <= 0 {
		return nil, errors.New("invalid permission id")
	}

	return s.PermissionStore.GetByID(
		ctx,
		id,
	)
}

// ============================================================
// GET BY USER ID
// ============================================================

func (s *PermissionService) GetPermissionByUserID(
	ctx context.Context,
	userID int64,
) ([]*permissiondto.Permission, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	return s.PermissionStore.GetPermissionByUserID(
		ctx,
		userID,
	)
}

// ============================================================
// GET BY ROLE ID
// ============================================================

func (s *PermissionService) GetByRoleID(
	ctx context.Context,
	roleID int64,
) ([]*permissiondto.Permission, error) {

	if roleID <= 0 {
		return nil, errors.New("invalid role id")
	}

	return s.PermissionStore.GetByRoleID(
		ctx,
		roleID,
	)
}

// Old code -------------------------> starts

// func (s *PermissionService) GetByID(
// 	ctx context.Context,
// 	id int64,
// ) (*permissiondto.Permission, error) {

// 	permission, err := s.PermissionStore.GetByID(
// 		ctx,
// 		id,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return permission, nil
// }
