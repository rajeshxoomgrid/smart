package internalsetup

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
	deptservice "github.com/rajeshbond/smart/internal/http/department_master/dept_service"
	"github.com/rajeshbond/smart/internal/http/tenant"
	userrole "github.com/rajeshbond/smart/internal/http/user_role"
	"github.com/rajeshbond/smart/internal/http/users"
)

type Service struct {
	db            *sql.DB
	tenantService *tenant.Service
	deptService   *deptservice.DeptService
	roleService   *userrole.Service
	userService   *users.Service
}

func NewService(
	db *sql.DB,
	tenantService *tenant.Service,
	deptService *deptservice.DeptService,
	roleService *userrole.Service,
	userService *users.Service,
) *Service {
	return &Service{
		db:            db,
		tenantService: tenantService,
		deptService:   deptService,
		roleService:   roleService,
		userService:   userService,
	}
}

func (s *Service) SetupSuperAdmin(
	ctx context.Context,
	dto *SetupSuperAdminDTO,
) (*SetupResponse, error) {

	// ============================================================
	// 1. VALIDATION
	// ============================================================

	if dto == nil {
		return nil, fmt.Errorf("setup request is required")
	}

	if s.db == nil {
		return nil, fmt.Errorf("database connection is required")
	}

	if s.roleService == nil ||
		s.tenantService == nil ||
		s.userService == nil ||
		s.deptService == nil {
		return nil, fmt.Errorf("setup service dependencies are not initialized")
	}

	role := strings.TrimSpace(dto.Role.UserRole)
	if role == "" {
		return nil, fmt.Errorf("user role is required")
	}

	tenantName := strings.TrimSpace(dto.Tenant.TenantName)
	if tenantName == "" {
		return nil, fmt.Errorf("tenant name is required")
	}

	tenantCode := strings.TrimSpace(dto.Tenant.TenantCode)
	if tenantCode == "" {
		return nil, fmt.Errorf("tenant code is required")
	}

	departmentName := strings.TrimSpace(dto.Department.Department)
	if departmentName == "" {
		return nil, fmt.Errorf("department is required")
	}

	if len(departmentName) > 100 {
		return nil, fmt.Errorf("department must not exceed 100 characters")
	}

	employeeID := strings.TrimSpace(dto.User.EmployeeID)
	if employeeID == "" {
		return nil, fmt.Errorf("employee ID is required")
	}

	userName := strings.TrimSpace(dto.User.UserName)
	if userName == "" {
		return nil, fmt.Errorf("username is required")
	}

	if dto.User.Password == "" {
		return nil, fmt.Errorf("password is required")
	}

	// ============================================================
	// 2. BEGIN TRANSACTION
	// ============================================================

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin setup transaction: %w", err)
	}

	committed := false

	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	// ============================================================
	// 3. CREATE ROLE
	// ============================================================

	roleID, err := s.roleService.CreateRoleTx(
		ctx,
		tx,
		role,
	)
	if err != nil {
		return nil, fmt.Errorf("create role: %w", err)
	}

	// ============================================================
	// 4. CREATE TENANT
	// ============================================================

	tenantReq := tenant.CreateTenantRequied(dto.Tenant)

	tenantReq.TenantName = tenantName
	tenantReq.TenantCode = tenantCode

	tenantID, err := s.tenantService.CreateSuperTenantTx(
		ctx,
		tx,
		tenantReq,
	)
	if err != nil {
		return nil, fmt.Errorf("create tenant: %w", err)
	}

	// ============================================================
	// 5. CREATE DEPARTMENT
	// ============================================================

	departmentReq := &deptdto.CreateDepartmentRequest{
		Department: departmentName,
	}

	department, err := s.deptService.CreateDepartmentTx(
		ctx,
		tx,
		1, // Temporary system creator ID; see note below.
		tenantID,
		departmentReq,
	)
	if err != nil {
		return nil, fmt.Errorf("create department: %w", err)
	}

	deptID := department.ID

	// ============================================================
	// 6. PREPARE USER REQUEST
	// ============================================================

	var phonePtr *string
	var emailPtr *string

	phone := strings.TrimSpace(dto.User.Phone)
	email := strings.TrimSpace(dto.User.Email)

	if phone != "" {
		phonePtr = &phone
	}

	if email != "" {
		emailPtr = &email
	}

	userReq := users.UserSuperRequest{
		EmployeeID: employeeID,
		UserName:   userName,
		Phone:      phonePtr,
		Email:      emailPtr,
		Password:   dto.User.Password,
	}

	// ============================================================
	// 7. CREATE SUPER ADMIN USER
	// ============================================================

	userID, err := s.userService.CreateSuperUserTx(
		ctx,
		tx,
		tenantID,
		roleID,
		deptID,
		userReq,
	)
	if err != nil {
		return nil, fmt.Errorf("create super admin user: %w", err)
	}

	// ============================================================
	// 8. COMMIT TRANSACTION
	// ============================================================

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit setup transaction: %w", err)
	}

	committed = true

	// ============================================================
	// 9. RETURN RESPONSE
	// ============================================================

	return &SetupResponse{
		RoleID:   roleID,
		TenantID: tenantID,
		UserID:   userID,
		DeptID:   deptID,
		Message:  "Successfully",
	}, nil
}

// func (s *Service) SetupSuperAdmin(ctx context.Context, dto *SetupSuperAdminDTO) (*SetupResponse, error) {
// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer tx.Rollback()

// 	// Create Role
// 	roleID, err := s.roleService.CreateRoleTx(ctx, tx, dto.Role.UserRole)

// 	// Create Tenant
// 	tenantID, err := s.tenantService.CreateSuperTenantTx(ctx, tx, tenant.CreateTenantRequied(dto.Tenant))

// 	// Create User
// 	var phonePtr *string
// 	var emailPtr *string

// 	if dto.User.Phone != "" {
// 		phonePtr = &dto.User.Phone
// 	}

// 	if dto.User.Email != "" {
// 		emailPtr = &dto.User.Email
// 	}

// 	userreq := users.UserSuperRequest{
// 		EmployeeID: dto.User.EmployeeID,
// 		UserName:   dto.User.UserName,
// 		Phone:      phonePtr,
// 		Email:      emailPtr,
// 		Password:   dto.User.Password,
// 	}
// 	userID, err := s.userService.CreateSuperUserTx(ctx, tx, tenantID, roleID, userreq)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &SetupResponse{
// 		RoleID:   roleID,
// 		TenantID: tenantID,
// 		UserID:   userID,
// 		Message:  "Sucessfully",
// 	}, nil

// }
