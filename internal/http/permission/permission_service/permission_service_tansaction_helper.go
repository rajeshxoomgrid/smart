package permissionservice

// func (s *PermissionService) withTransaction(
// 	ctx context.Context,
// 	fn func(tx *sql.Tx) error,
// ) error {

// 	tx, err := s.PermissionStore.BeginTx(
// 		ctx,
// 		nil,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		_ = tx.Rollback()
// 	}()

// 	if err := fn(tx); err != nil {
// 		return err
// 	}

// 	return tx.Commit()
// }
