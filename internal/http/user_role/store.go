package userrole

// Index
// 1. Create
// 2. GetRoleIDByName
// 3. GetRoleNameByID
// 4. CreateRoleSuperTx
// 5. Role present in DB

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Function Starts here -------

// 1. Create

func (s *Store) Create(
	ctx context.Context,
	userRole string,
	tenantID int64,
	userID int64,
) (*UserRoleResponseDTO, error) {

	userRole = strings.ToLower(strings.TrimSpace(userRole))

	if userRole == "" {
		return nil, fmt.Errorf("user role is required")
	}

	query := `
		INSERT INTO user_role (
			user_role,
			tenant_id,
			created_by,
			updated_by
		)
		VALUES ($1, $2, $3, $3)
		ON CONFLICT (user_role, tenant_id) DO NOTHING
		RETURNING
			id,
			user_role,
			tenant_id,
			created_by,
			updated_by,
			created_at,
			updated_at
	`

	var role UserRoleResponseDTO

	err := s.db.QueryRowContext(
		ctx,
		query,
		userRole,
		tenantID,
		userID,
	).Scan(
		&role.ID,
		&role.UserRole,
		&role.TenantID,
		&role.CreatedBy,
		&role.UpdatedBy,
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role already exists for this tenant")
		}

		return nil, err
	}

	return &role, nil
}

// func (s *Store) Create(
// 	ctx context.Context,
// 	user_role string,
// 	tenantID int64,
// 	userID int64,
// ) (*UserRoleResponseDTO, error) {

// 	user_role = strings.ToLower(strings.TrimSpace(user_role))

// 	if user_role == "" {
// 		return nil, fmt.Errorf("user role is required")
// 	}

// 	query := `
// 		INSERT INTO user_role (
// 			user_role,
// 			tenant_id,
// 			created_by,
// 			updated_by
// 		)
// 		VALUES ($1, $2, $3, $3)
// 		ON CONFLICT (user_role) DO NOTHING
// 		RETURNING
// 			id,
// 			user_role,
// 			tenant_id,
// 			created_by,
// 			updated_by,
// 			created_at,
// 			updated_at
// 	`

// 	var role UserRoleResponseDTO

// 	err := s.db.QueryRowContext(
// 		ctx,
// 		query,
// 		user_role,
// 		tenantID,
// 		userID,
// 	).Scan(
// 		&role.ID,
// 		&role.UserRole,
// 		&role.TenantID,
// 		&role.CreatedBy,
// 		&role.UpdatedBy,
// 		&role.CreatedAt,
// 		&role.UpdatedAt,
// 	)

// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, fmt.Errorf("role already exists")
// 		}

// 		return nil, err
// 	}

// 	return &role, nil
// }

// func (s *Store) Create(ctx context.Context, user_role string, tenantID int64, userID int64) (*UserRoleResponseDTO, error) {

// 	query := `
// 	INSERT INTO user_role (user_role,tenant_id ,created_by, updated_by)
// 	VALUES ($1,$2 $3, $3)
// 	ON CONFLICT (user_role) DO NOTHING
// 	RETURNING id, user_role,tenant_id ,created_by, updated_by, created_at, updated_at
// 	`

// 	role := UserRoleResponseDTO
// 	err := s.db.QueryRowContext(
// 		ctx,
// 		query,
// 		strings.ToLower(user_role),
// 		tenantID,
// 		userID).
// 		Scan(&role.ID, &role.UserRole, &role.TenantID, &role.CreatedBy, &role.UpdatedBy, &role.CreatedAt, &role.UpdatedAt)

// 	if err != nil {

// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("role already exists ")
// 		}

// 		return nil, err
// 	}

// 	return role, nil
// }

// 2. GetRoleIDByName
func (s *Store) GetRoleIDByName(ctx context.Context, roleName string) (int64, error) {
	query := `
	SELECT id FROM user_role
	WHERE LOWER(user_role) = LOWER($1)
	`

	var roleID int64

	err := s.db.QueryRowContext(
		ctx,
		query,
		strings.ToLower(roleName)).Scan(&roleID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("Invalid role")
		}
		return 0, err
	}

	return roleID, nil

}

// 3. GetRoleNameByID
func (s *Store) GetRoleNameByID(ctx context.Context, roleID int64) (string, error) {

	query := `

	SELECT user_role
	FROM user_role
	WHERE id = $1
	`

	var roleName string

	err := s.db.QueryRowContext(
		ctx,
		query,
		roleID,
	).Scan(&roleName)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("role not found")
		}
		return "", nil
	}
	return roleName, nil
}

// 4. CreateRoleSuperTx - (only for development phase	)

func (s *Store) CreateRoleSuperTx(
	ctx context.Context,
	tx *sql.Tx,
	tenant_id int64,
	dto CreateUserRoleDTO,
) (int64, error) {

	query := `
        INSERT INTO user_role (
            tenant_id,
            user_role,
            created_by,
            updated_by
        )
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	var roleID int64

	err := tx.QueryRowContext(
		ctx,
		query,
		tenant_id,
		strings.ToLower(strings.TrimSpace(dto.UserRole)),
		dto.CreatedBy,
		dto.UpdatedBy,
	).Scan(&roleID)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return 0, errors.New("role already exists")
			}

			if pqErr.Code == "23503" {
				return 0, errors.New("invalid tenant ID")
			}
		}

		return 0, err
	}

	return roleID, nil
}

// func (s *Store) CreateRoleSuperTx(ctx context.Context, tx *sql.Tx,tenant_id int64,dto CreateUserRoleDTO) (int64, error) {
// 	query := `
// 		INSERT INTO "user_role" (tenant_id,user_role,created_by,updated_by)
// 		VALUES($1,$2,$3,$4)
// 		RETURNING id
// 	`
// 	var roleID int64

// 	err := tx.QueryRowContext(
// 		ctx,
// 		query,
// 		tenant_id,
// 		strings.ToLower(dto.UserRole),
// 		dto.CreatedBy,
// 		dto.UpdatedBy,
// 	).Scan(&roleID)

// 	if err != nil {

// 		// Handle duplicate key error
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			if pqErr.Code == "23505" {
// 				return 0, errors.New("role already exists")
// 			}
// 		}

// 		return 0, err

// 	}

// 	return roleID, nil

// }

// 5. Role present in DB

func (s *Store) RoleInDB(ctx context.Context, roleName string) (bool, error) {
	query := `
	SELECT EXISTS(
		SELECT 1
		FROM user_role
		WHERE LOWER(user_role) = LOWER($1)
	)
	`

	var exists bool
	err := s.db.QueryRowContext(ctx, query, roleName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
