package permissiondto

import "time"

// Permission represents a permission record.
type Permission struct {
	ID int64 `json:"id"`

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

// CreatePermissionRequest represents the create request.
type CreatePermissionRequest struct {
	AllPerm        bool `json:"all_perm"`
	CreatePerm     bool `json:"create_perm"`
	ReadPerm       bool `json:"read_perm"`
	UpdatePerm     bool `json:"update_perm"`
	TempUpdatePerm bool `json:"temp_update_permission"`
	DeletePerm     bool `json:"delete_perm"`

	CreatedBy *int64 `json:"created_by,omitempty"`
}

// UpdatePermissionRequest represents the update request.
type UpdatePermissionRequest struct {
	AllPerm    bool `json:"all_perm"`
	CreatePerm bool `json:"create_perm"`
	ReadPerm   bool `json:"read_perm"`
	UpdatePerm bool `json:"update_perm"`
	DeletePerm bool `json:"delete_perm"`

	UpdatedBy *int64 `json:"updated_by,omitempty"`
}

// PermissionResponse represents a single permission response.
type PermissionResponse struct {
	Data *Permission `json:"data"`
}

// PermissionListResponse represents a list of permissions.
type PermissionListResponse struct {
	Data  []*Permission `json:"data"`
	Total int           `json:"total"`
}

// PermissionFilter represents a permission filter.
type PermissionFilter struct {
	Page      int    `json:"page" query:"page"`
	PageSize  int    `json:"page_size" query:"page_size"`
	Search    string `json:"search" query:"search"`
	SortBy    string `json:"sort_by" query:"sort_by"`
	SortOrder string `json:"sort_order" query:"sort_order"`
}
