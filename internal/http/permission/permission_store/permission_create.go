package permissionstore

import (
	"context"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

// ============================================================
// CREATE PERMISSION
// ============================================================

func (s *PermissionStore) Create(
	ctx context.Context,
	req *permissiondto.CreatePermissionRequest,
) (*permissiondto.Permission, error) {

	query := `
		INSERT INTO public.permission (
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
		VALUES (
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

	p := &permissiondto.Permission{}

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
		&p.ID,
		&p.UserID,
		&p.RoleID,
		&p.DepartmentID,
		&p.AllPerm,
		&p.CreatePerm,
		&p.ReadPerm,
		&p.UpdatePerm,
		&p.TempUpdatePerm,
		&p.DeletePerm,
		&p.CreatedBy,
		&p.UpdatedBy,
		&p.DeletedBy,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
		&p.IsDeleted,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}
