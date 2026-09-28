package permissionservice

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/users"
)

type UserProvider interface {
	CheckEmployeeExist(ctx context.Context, employeeID string, tenantID int64) error

	GetUserByEmployeeID(ctx context.Context, employeeID string, tenantID int64) (*users.UserResponse, error)
}
