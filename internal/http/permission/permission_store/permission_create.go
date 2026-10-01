package permissionstore

import (
	"context"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (s *PermissionStore) Create(ctx context.Context, req *permissiondto.CreatePermissionRequest) (*permissiondto.Permission, error) {

	query := `
		INSERT INTO public.permission
		(
			user_id,
			role_id,
			department_id,
			all_perm,
			create_perm,
			read_perm,
			update_perm,
			temp_update_permission,
			delete_perm,
			created_by
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10
		)
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
		req.UserID,
		req.RoleID,
		req.DepartmentID,
		req.AllPerm,
		req.CreatePerm,
		req.ReadPerm,
		req.UpdatePerm,
		req.TempUpdatePerm,
		req.DeletePerm,
		req.CreatedBy,
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

