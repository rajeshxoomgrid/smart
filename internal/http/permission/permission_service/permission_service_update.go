package permissionservice

import (
	"context"
	"errors"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionService) Update(
	ctx context.Context,
	id int64,
	req *permissiondto.UpdatePermissionRequest,
) (*permissiondto.PermissionResponse, error) {

	// --------------------------------------------------
	// Validate ID
	// --------------------------------------------------

	if id <= 0 {
		return nil, errors.New("invalid permission id")
	}

	// --------------------------------------------------
	// Validate request
	// --------------------------------------------------

	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	// --------------------------------------------------
	// Validate UpdatedBy
	// --------------------------------------------------

	if req.UpdatedBy == nil || *req.UpdatedBy <= 0 {
		return nil, errors.New("invalid updated_by")
	}

	// --------------------------------------------------
	// Store
	// --------------------------------------------------

	permission, err := s.PermissionStore.Update(
		ctx,
		id,
		req,
	)

	if err != nil {
		return nil, err
	}

	return &permissiondto.PermissionResponse{
		Data: permission,
	}, nil
}

// func (s *PermissionService) Update(
// 	ctx context.Context,
// 	id int64,
// 	req *permissiondto.UpdatePermissionRequest,
// ) (*permissiondto.Permission, error) {

// 	var permission *permissiondto.Permission

// 	err := s.withTransaction(
// 		ctx,
// 		func(tx *sql.Tx) error {

// 			var err error

// 			permission, err = s.PermissionStore.Update(
// 				ctx,
// 				tx,
// 				id,
// 				req,
// 			)

// 			if err != nil {
// 				return err
// 			}

// 			return nil
// 		},
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return permission, nil
// }
