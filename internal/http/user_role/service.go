package userrole

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/auth"
)

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) Create(
	ctx context.Context,
	dto CreateRole, tenantID, userID int64,
) (*UserRoleResponseDTO, error) {

	fmt.Println("Service", dto)

	sendrole := dto.UserRole

	if sendrole == "" {
		return nil, fmt.Errorf("user_role is required")
	}

	exist, err := s.store.GetRoleIDByName(ctx, sendrole, tenantID)

	if err != nil {
		return nil, err
	}

	if exist <= 0 {
		return nil, fmt.Errorf("role already exists")
	}

	// role, err := s.store.Create(ctx, dto)
	role, err := s.store.Create(ctx, dto.UserRole, tenantID, userID)
	if err != nil {
		return nil, err
	}

	return &UserRoleResponseDTO{
		ID:        role.ID,
		TenantID:  role.TenantID,
		UserRole:  role.UserRole,
		CreatedBy: role.CreatedBy,
		UpdatedBy: role.UpdatedBy,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}, nil
}

// This is for intial route only allowed in developemt phase
func (s *Service) CreateRoleTx(ctx context.Context, tx *sql.Tx, tenant_id int64, role string) (int64, error) {

	createdBy := int64(1)

	dto := CreateUserRoleDTO{
		UserRole:  strings.ToLower(role),
		CreatedBy: &createdBy,
		UpdatedBy: &createdBy,
	}

	return s.store.CreateRoleSuperTx(ctx, tx, tenant_id, dto)

}

func (s *Service) CreateRoleBySuper(
	ctx context.Context,
	req CreateRoleBySuper,
	claims *auth.UserClaims,
) (*UserRoleResponseDTO, error) {

	if req.TenantID <= 0 {
		return nil, fmt.Errorf("tenant ID is required")
	}

	userRole := strings.TrimSpace(req.UserRole)

	if userRole == "" {
		return nil, fmt.Errorf("user role is required")
	}

	return s.store.Create(
		ctx,
		strings.ToLower(userRole),
		req.TenantID,
		claims.UserID,
	)
}
func (s *Service) GetRoleNameByID(ctx context.Context, roleID int64) (string, error) {

	return s.store.GetRoleNameByID(ctx, roleID)
}

// Get the Role id from user role

func (s *Service) GetRoleIDByName(ctx context.Context, roleName string, tenantID int64) (int64, error) {

	return s.store.GetRoleIDByName(ctx, roleName, tenantID)
}

// Delete the user-role

func (s *Service) DeleteUserrole(ctx context.Context, roleID, tenatID int64) (bool, error) {
	if roleID <= 0 {
		return false, fmt.Errorf("Role ID should not zero")
	}

	if tenatID <= 0 {
		return false, fmt.Errorf("Tenant ID should not zero")
	}
	return s.store.DeleteUserRole(ctx, roleID, tenatID)
}

// Get the UserName by ID

func (s *Service) GetUserRoleIDByName(ctx context.Context, roleID, tenantID int64) (string, error) {
	if roleID <= 0 {
		return "", fmt.Errorf("Role ID should not zero")
	}

	if tenantID <= 0 {
		return "", fmt.Errorf("Tenant ID should not zero")
	}

	return s.store.GetTheUserNameByID(ctx, roleID, tenantID)
}
