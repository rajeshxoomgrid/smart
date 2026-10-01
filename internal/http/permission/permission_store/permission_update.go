package permissionstore

import (
	"context"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

// ============================================================
// UPDATE
// ============================================================

func (s *PermissionStore) Update(
	ctx context.Context,
	id int64,
	req *permissiondto.UpdatePermissionRequest,
) (*permissiondto.Permission, error) {

	query := `
		UPDATE public.permission
		SET
			user_id = $2,
			role_id = $3,
			department_id = $4,
			all_perm = $5,
			create_perm = $6,
			read_perm = $7,
			update_perm = $8,
			temp_update_permission = $9,
			delete_perm = $10,
			updated_by = $11,
			updated_at = NOW()
		WHERE id = $1
		  AND is_deleted = false
		RETURNING
			id,
			user_id,
			role_id,
			department_id,
			all_perm,
			create_perm,
			read_perm,
			update_perm,
			temp_update_permission,
			delete_perm,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at,
			is_deleted
	`

	result := &permissiondto.Permission{}

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
		req.UserID,
		req.RoleID,
		req.DepartmentID,
		req.AllPerm,
		req.CreatePerm,
		req.ReadPerm,
		req.UpdatePerm,
		req.TempUpdatePerm,
		req.DeletePerm,
		req.UpdatedBy,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.RoleID,
		&result.DepartmentID,
		&result.AllPerm,
		&result.CreatePerm,
		&result.ReadPerm,
		&result.UpdatePerm,
		&result.TempUpdatePerm,
		&result.DeletePerm,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.DeletedBy,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.IsDeleted,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// old code

// func (s *PermissionStore) update(
// 	ctx context.Context,
// 	tx *sql.Tx,
// 	id int64,
// 	req *permissiondto.UpdatePermissionRequest,
// ) (*permissiondto.Permission, error) {

// 	var permission permissiondto.Permission

// 	err := tx.QueryRowContext(
// 		ctx,
// 		queryUpdatePermission,
// 		req.AllPerm,
// 		req.CreatePerm,
// 		req.ReadPerm,
// 		req.UpdatePerm,
// 		req.DeletePerm,
// 		req.UpdatedBy,
// 		id,
// 	).Scan(
// 		&permission.ID,
// 		&permission.AllPerm,
// 		&permission.CreatePerm,
// 		&permission.ReadPerm,
// 		&permission.UpdatePerm,
// 		&permission.DeletePerm,
// 		&permission.CreatedBy,
// 		&permission.UpdatedBy,
// 		&permission.DeletedBy,
// 		&permission.CreatedAt,
// 		&permission.UpdatedAt,
// 		&permission.DeletedAt,
// 		&permission.IsDeleted,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &permission, nil
// }
