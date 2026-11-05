package deptstore

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

// ============================================================
// LIST
// ============================================================

func (s *DeptStore) List(
	ctx context.Context,
	tx *sql.Tx,
	req *deptdto.ListRequest,
) ([]deptdto.Department, int64, error) {

	offset := (req.Page - 1) * req.Limit

	search := strings.TrimSpace(
		req.Search,
	)

	// ============================================================
	// COUNT
	// ============================================================

	countQuery := queryCountDepartment

	countArgs := []interface{}{}

	argIndex := 1

	if search != "" {

		countQuery += `
            AND department ILIKE $` +
			fmt.Sprint(argIndex)

		countArgs = append(
			countArgs,
			"%"+search+"%",
		)

		argIndex++
	}

	var total int64

	err := tx.QueryRowContext(
		ctx,
		countQuery,
		countArgs...,
	).Scan(&total)

	if err != nil {
		return nil, 0, fmt.Errorf(
			"count departments: %w",
			err,
		)
	}

	// ============================================================
	// DATA QUERY
	// ============================================================

	query := queryListDepartment

	args := []interface{}{}

	argIndex = 1

	if search != "" {

		query += `
            AND department ILIKE $` +
			fmt.Sprint(argIndex)

		args = append(
			args,
			"%"+search+"%",
		)

		argIndex++
	}

	// ============================================================
	// PAGINATION
	// ============================================================

	query += `
        ORDER BY id DESC
        LIMIT $` + fmt.Sprint(argIndex) + `
        OFFSET $` + fmt.Sprint(argIndex+1)

	args = append(
		args,
		req.Limit,
		offset,
	)

	// ============================================================
	// EXECUTE
	// ============================================================

	rows, err := tx.QueryContext(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, 0, fmt.Errorf(
			"list departments: %w",
			err,
		)
	}

	defer rows.Close()

	// ============================================================
	// SCAN
	// ============================================================

	items := make(
		[]deptdto.Department,
		0,
	)

	for rows.Next() {

		var item deptdto.Department

		err := rows.Scan(
			&item.ID,
			&item.Department,
			&item.CreatedBy,
			&item.UpdatedBy,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.IsDeleted,
			&item.DeletedAt,
		)

		if err != nil {
			return nil, 0, fmt.Errorf(
				"scan department: %w",
				err,
			)
		}

		items = append(
			items,
			item,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf(
			"iterate departments: %w",
			err,
		)
	}

	return items, total, nil
}
