package permissionstore

import (
	"context"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

// ============================================================
// GET BY ID
// ============================================================

func (s *PermissionStore) GetByID(
	ctx context.Context,
	id int64,
) (*permissiondto.Permission, error) {

	query := `
		SELECT
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
		FROM public.permission
		WHERE id = $1
		  AND is_deleted = false
	`

	result := &permissiondto.Permission{}

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
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

// ============================================================
// GET BY USER ID
// ============================================================

func (s *PermissionStore) GetByUserID(
	ctx context.Context,
	userID int64,
) ([]*permissiondto.Permission, error) {

	query := `
		SELECT
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
		FROM public.permission
		WHERE user_id = $1
		  AND is_deleted = false
		ORDER BY id DESC
	`

	rows, err := s.db.QueryContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*permissiondto.Permission, 0)

	for rows.Next() {

		item := &permissiondto.Permission{}

		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.RoleID,
			&item.DepartmentID,
			&item.AllPerm,
			&item.CreatePerm,
			&item.ReadPerm,
			&item.UpdatePerm,
			&item.TempUpdatePerm,
			&item.DeletePerm,
			&item.CreatedBy,
			&item.UpdatedBy,
			&item.DeletedBy,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
			&item.IsDeleted,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// ============================================================
// GET BY ROLE ID
// ============================================================

func (s *PermissionStore) GetByRoleID(
	ctx context.Context,
	roleID int64,
) ([]*permissiondto.Permission, error) {

	query := `
		SELECT
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
		FROM public.permission
		WHERE role_id = $1
		  AND is_deleted = false
		ORDER BY id DESC
	`

	return s.getMany(ctx, query, roleID)
}

// ============================================================
// GET BY DEPARTMENT ID
// ============================================================

func (s *PermissionStore) GetByDepartmentID(
	ctx context.Context,
	departmentID int64,
) ([]*permissiondto.Permission, error) {

	query := `
		SELECT
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
		FROM public.permission
		WHERE department_id = $1
		  AND is_deleted = false
		ORDER BY id DESC
	`

	return s.getMany(ctx, query, departmentID)
}

// ============================================================
// COMMON GET MANY
// ============================================================

func (s *PermissionStore) getMany(
	ctx context.Context,
	query string,
	id int64,
) ([]*permissiondto.Permission, error) {

	rows, err := s.db.QueryContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*permissiondto.Permission, 0)

	for rows.Next() {

		item := &permissiondto.Permission{}

		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.RoleID,
			&item.DepartmentID,
			&item.AllPerm,
			&item.CreatePerm,
			&item.ReadPerm,
			&item.UpdatePerm,
			&item.TempUpdatePerm,
			&item.DeletePerm,
			&item.CreatedBy,
			&item.UpdatedBy,
			&item.DeletedBy,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
			&item.IsDeleted,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
