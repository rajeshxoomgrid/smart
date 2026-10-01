package permissionstore

import (
	"context"
	"fmt"
	"strings"

	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

// func (s *PermissionStore) list(
// 	ctx context.Context,
// 	db *sql.DB,
// 	filter *permissiondto.PermissionFilter,
// ) ([]*permissiondto.Permission, int, error) {

// 	if filter.Page <= 0 {
// 		filter.Page = 1
// 	}

// 	if filter.PageSize <= 0 {
// 		filter.PageSize = 20
// 	}

// 	if filter.PageSize > 100 {
// 		filter.PageSize = 100
// 	}

// 	offset := (filter.Page - 1) * filter.PageSize

// 	args := []any{}
// 	where := []string{
// 		"is_deleted = FALSE",
// 	}

// 	if filter.Search != "" {
// 		args = append(args, "%"+filter.Search+"%")

// 		where = append(
// 			where,
// 			fmt.Sprintf(
// 				"CAST(id AS TEXT) ILIKE $%d",
// 				len(args),
// 			),
// 		)
// 	}

// 	whereClause := strings.Join(where, " AND ")

// 	countQuery := fmt.Sprintf(`
// 		SELECT COUNT(*)
// 		FROM permission_master
// 		WHERE %s
// 	`, whereClause)

// 	var total int

// 	err := db.QueryRowContext(
// 		ctx,
// 		countQuery,
// 		args...,
// 	).Scan(&total)

// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	sortBy := "id"

// 	switch filter.SortBy {
// 	case "id":
// 		sortBy = "id"
// 	case "created_at":
// 		sortBy = "created_at"
// 	case "updated_at":
// 		sortBy = "updated_at"
// 	}

// 	sortOrder := "DESC"

// 	if strings.EqualFold(filter.SortOrder, "asc") {
// 		sortOrder = "ASC"
// 	}

// 	args = append(args, filter.PageSize)
// 	limitPosition := len(args)

// 	args = append(args, offset)
// 	offsetPosition := len(args)

// 	listQuery := fmt.Sprintf(`
// 		SELECT
// 			id,
// 			all_perm,
// 			create_perm,
// 			read_perm,
// 			update_perm,
// 			delete_perm,
// 			created_by,
// 			updated_by,
// 			deleted_by,
// 			created_at,
// 			updated_at,
// 			deleted_at,
// 			is_deleted
// 		FROM permission_master
// 		WHERE %s
// 		ORDER BY %s %s
// 		LIMIT $%d
// 		OFFSET $%d
// 	`, whereClause, sortBy, sortOrder, limitPosition, offsetPosition)

// 	rows, err := db.QueryContext(
// 		ctx,
// 		listQuery,
// 		args...,
// 	)

// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	defer rows.Close()

// 	permissions := make([]*permissiondto.Permission, 0)

// 	for rows.Next() {

// 		var permission permissiondto.Permission

// 		err := rows.Scan(
// 			&permission.ID,
// 			&permission.AllPerm,
// 			&permission.CreatePerm,
// 			&permission.ReadPerm,
// 			&permission.UpdatePerm,
// 			&permission.DeletePerm,
// 			&permission.CreatedBy,
// 			&permission.UpdatedBy,
// 			&permission.DeletedBy,
// 			&permission.CreatedAt,
// 			&permission.UpdatedAt,
// 			&permission.DeletedAt,
// 			&permission.IsDeleted,
// 		)

// 		if err != nil {
// 			return nil, 0, err
// 		}

// 		permissions = append(
// 			permissions,
// 			&permission,
// 		)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, 0, err
// 	}

// 	return permissions, total, nil
// }

// ============================================================
// LIST
// ============================================================

func (s *PermissionStore) List(
	ctx context.Context,
	filter *permissiondto.PermissionFilter,
) ([]*permissiondto.Permission, int, error) {

	page := filter.Page

	if page <= 0 {
		page = 1
	}

	pageSize := filter.PageSize

	if pageSize <= 0 {
		pageSize = 20
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// --------------------------------------------------------
	// SORT COLUMN
	// --------------------------------------------------------

	sortColumn := "p.id"

	switch strings.ToLower(filter.SortBy) {

	case "id":
		sortColumn = "p.id"

	case "user_id":
		sortColumn = "p.user_id"

	case "role_id":
		sortColumn = "p.role_id"

	case "department_id":
		sortColumn = "p.department_id"

	case "created_at":
		sortColumn = "p.created_at"
	}

	// --------------------------------------------------------
	// SORT ORDER
	// --------------------------------------------------------

	sortOrder := "DESC"

	if strings.ToUpper(filter.SortOrder) == "ASC" {
		sortOrder = "ASC"
	}

	// --------------------------------------------------------
	// WHERE
	// --------------------------------------------------------

	where := `
		WHERE p.is_deleted = false
	`

	args := make([]interface{}, 0)

	argIndex := 1

	search := strings.TrimSpace(filter.Search)

	if search != "" {

		where += fmt.Sprintf(`
			AND (
				CAST(p.id AS TEXT) ILIKE $%d
				OR CAST(p.user_id AS TEXT) ILIKE $%d
				OR CAST(p.role_id AS TEXT) ILIKE $%d
				OR CAST(p.department_id AS TEXT) ILIKE $%d
			)
		`,
			argIndex,
			argIndex,
			argIndex,
			argIndex,
		)

		args = append(
			args,
			"%"+search+"%",
		)

		argIndex++
	}

	// --------------------------------------------------------
	// COUNT
	// --------------------------------------------------------

	countQuery := `
		SELECT COUNT(*)
		FROM public.permission p
	` + where

	var total int

	err := s.db.QueryRowContext(
		ctx,
		countQuery,
		args...,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	// --------------------------------------------------------
	// DATA
	// --------------------------------------------------------

	dataQuery := fmt.Sprintf(`
		SELECT
			p.id,
			p.user_id,
			p.role_id,
			p.department_id,
			p.all_perm,
			p.create_perm,
			p.read_perm,
			p.update_perm,
			p.temp_update_permission,
			p.delete_perm,
			p.created_by,
			p.updated_by,
			p.deleted_by,
			p.created_at,
			p.updated_at,
			p.deleted_at,
			p.is_deleted
		FROM public.permission p
		%s
		ORDER BY %s %s
		LIMIT $%d
		OFFSET $%d
	`,
		where,
		sortColumn,
		sortOrder,
		argIndex,
		argIndex+1,
	)

	dataArgs := append([]interface{}{}, args...)

	dataArgs = append(
		dataArgs,
		pageSize,
		offset,
	)

	rows, err := s.db.QueryContext(
		ctx,
		dataQuery,
		dataArgs...,
	)

	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}

		result = append(
			result,
			item,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return result, total, nil
}
