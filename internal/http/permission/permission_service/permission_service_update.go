package permissionservice

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
