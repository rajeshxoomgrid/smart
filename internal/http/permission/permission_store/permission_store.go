package permissionstore

import (
	"database/sql"
)

type PermissionStore struct {
	db *sql.DB
}

func NewPermissionStore(db *sql.DB) *PermissionStore {
	return &PermissionStore{db: db}
}

// Old code-------------------------> starts

// Create creates a permission using transaction.
// func (s *PermissionStore) Create(
// 	ctx context.Context,
// 	tx *sql.Tx,
// 	req *permissiondto.CreatePermissionRequest,
// ) (*permissiondto.Permission, error) {
// 	return s.CreatePermission(ctx, tx, req)
// }

// // GetByID returns a permission by ID.
// func (s *PermissionStore) GetByID(
// 	ctx context.Context,
// 	id int64,
// ) (*permissiondto.Permission, error) {
// 	return s.getByID(ctx, s.db, id)
// }

// List returns permissions.
// func (s *PermissionStore) List(
// 	ctx context.Context,
// 	filter *permissiondto.PermissionFilter,
// ) ([]*permissiondto.Permission, int, error) {
// 	return s.list(ctx, s.db, filter)
// }

// Update updates a permission using transaction.
// func (s *PermissionStore) Update(
// 	ctx context.Context,
// 	tx *sql.Tx,
// 	id int64,
// 	req *permissiondto.UpdatePermissionRequest,
// ) (*permissiondto.Permission, error) {
// 	return s.update(ctx, tx, id, req)
// }

// // Delete soft deletes a permission using transaction.
// func (s *PermissionStore) Delete(
// 	ctx context.Context,
// 	tx *sql.Tx,
// 	id int64,
// 	deletedBy *int64,
// ) error {
// 	return s.delete(ctx, tx, id, deletedBy)
// }

// // BeginTx starts a database transaction.
// func (s *PermissionStore) BeginTx(
// 	ctx context.Context,
// 	opts *sql.TxOptions,
// ) (*sql.Tx, error) {
// 	return s.db.BeginTx(ctx, opts)
// }
