package permissiondto

import "time"

// ============================================================
// Permission
// ============================================================

type Permission struct {
	ID int64 `json:"id"`

	UserID       int64 `json:"user_id"`
	RoleID       int64 `json:"role_id"`
	DepartmentID int64 `json:"dept_id"`

	AllPerm        bool `json:"all_perm"`
	CreatePerm     bool `json:"create_perm"`
	ReadPerm       bool `json:"read_perm"`
	UpdatePerm     bool `json:"update_perm"`
	TempUpdatePerm bool `json:"temp_update_permission"`
	DeletePerm     bool `json:"delete_perm"`

	CreatedBy *int64 `json:"created_by,omitempty"`
	UpdatedBy *int64 `json:"updated_by,omitempty"`
	DeletedBy *int64 `json:"deleted_by,omitempty"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	IsDeleted bool `json:"is_deleted"`
}

// ============================================================
// Create Permission Request
// ============================================================

type CreatePermissionRequest struct {
	UserID       int64 `json:"user_id" validate:"required"`
	RoleID       int64 `json:"role_id" validate:"required"`
	DepartmentID int64 `json:"dept_id" validate:"required"`

	AllPerm        bool `json:"all_perm"`
	CreatePerm     bool `json:"create_perm"`
	ReadPerm       bool `json:"read_perm"`
	UpdatePerm     bool `json:"update_perm"`
	TempUpdatePerm bool `json:"temp_update_permission"`
	DeletePerm     bool `json:"delete_perm"`

	CreatedBy *int64 `json:"created_by,omitempty"`
}

// ============================================================
// Update Permission Request
// ============================================================

type UpdatePermissionRequest struct {
	UserID       int64 `json:"user_id" validate:"required"`
	RoleID       int64 `json:"role_id" validate:"required"`
	DepartmentID int64 `json:"dept_id" validate:"required"`

	AllPerm        bool `json:"all_perm"`
	CreatePerm     bool `json:"create_perm"`
	ReadPerm       bool `json:"read_perm"`
	UpdatePerm     bool `json:"update_perm"`
	TempUpdatePerm bool `json:"temp_update_permission"`
	DeletePerm     bool `json:"delete_perm"`

	UpdatedBy *int64 `json:"updated_by,omitempty"`
}

// ============================================================
// Response
// ============================================================

type PermissionResponse struct {
	Data *Permission `json:"data"`
}

// ============================================================
// List Response
// ============================================================

type PermissionListResponse struct {
	Data  []*Permission `json:"data"`
	Total int           `json:"total"`
}

// ============================================================
// Filter
// ============================================================

type PermissionFilter struct {
	Page      int    `json:"page" query:"page"`
	PageSize  int    `json:"page_size" query:"page_size"`
	Search    string `json:"search" query:"search"`
	SortBy    string `json:"sort_by" query:"sort_by"`
	SortOrder string `json:"sort_order" query:"sort_order"`
}
