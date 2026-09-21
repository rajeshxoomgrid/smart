package deptdto

import "time"

// ============================================================
// Department
// ============================================================

type Department struct {
	ID         int64      `json:"id"`
	Department string     `json:"department"`
	CreatedBy  *int64     `json:"created_by,omitempty"`
	UpdatedBy  *int64     `json:"updated_by,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	IsDeleted  bool       `json:"is_deleted"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

// ============================================================
// Create Request
// ============================================================

type CreateDepartmentRequest struct {
	Department string `json:"department"`
}

// ============================================================
// Update Request
// ============================================================

type UpdateRequest struct {
	Department string `json:"department"`
}

// ============================================================
// List Request
// ============================================================

type ListRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

// ============================================================
// Pagination
// ============================================================

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// ============================================================
// List Response
// ============================================================

type ListResponse struct {
	Data       []Department `json:"data"`
	Pagination Pagination   `json:"pagination"`
}

// ============================================================
// Message Response
// ============================================================

type MessageResponse struct {
	Message string      `json:"message"`
	Data    *Department `json:"data,omitempty"`
}
