package partdto

import "time"

type CreatePartRequest struct {
	PartNo      string   `json:"part_no"`
	ProductName string   `json:"product_name"`
	Color       *string  `json:"color"`
	RMGrade     *string  `json:"rm_grade"`
	PartWt      *float64 `json:"part_wt"`
	Description *string  `json:"description"`
	IsActive    *bool    `json:"is_active"`
}

type UpdatePartRequest struct {
	PartNo      string   `json:"part_no"`
	ProductName string   `json:"product_name"`
	Color       *string  `json:"color"`
	RMGrade     *string  `json:"rm_grade"`
	PartWt      *float64 `json:"part_wt"`
	Description *string  `json:"description"`
	IsActive    bool     `json:"is_active"`
}

type PartResponse struct {
	ID          int64     `json:"id"`
	TenantID    int64     `json:"tenant_id"`
	PartNo      string    `json:"part_no"`
	ProductName string    `json:"product_name"`
	Color       *string   `json:"color"`
	RMGrade     *string   `json:"rm_grade"`
	PartWt      *float64  `json:"part_wt"`
	Description *string   `json:"description"`
	IsActive    bool      `json:"is_active"`
	IsDeleted   bool      `json:"is_deleted"`
	DeletedBy   *int64    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListPartsRequest struct {
	Search   string
	IsActive *bool
	Page     int
	PageSize int
}

type ListPartsResponse struct {
	Items      []PartResponse `json:"items"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	Total      int64          `json:"total"`
	TotalPages int64          `json:"total_pages"`
}
