package mold

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) BulkCreate(
	ctx context.Context,
	tx *sql.Tx,
	tenantID int64,
	userID int64,
	req BulkCreateMoldRequest,
) (inserted, skipped int, err error) {

	// -------------------------------------------------
	// 1. Validate request
	// -------------------------------------------------

	if len(req.Molds) == 0 {
		return 0, 0, nil
	}

	// -------------------------------------------------
	// 2. Collect mold numbers from request
	// -------------------------------------------------

	moldNos := make([]string, 0, len(req.Molds))

	for _, m := range req.Molds {
		moldNos = append(moldNos, m.MoldNo)
	}

	// -------------------------------------------------
	// 3. Fetch already existing molds
	//
	// Only check molds that:
	// - belong to the current tenant
	// - are not deleted
	// - have a mold_no present in the request
	// -------------------------------------------------

	checkQuery := `
		SELECT mold_no
		FROM mold
		WHERE tenant_id = $1
		  AND is_deleted = FALSE
		  AND mold_no = ANY($2)
	`

	rows, err := tx.QueryContext(
		ctx,
		checkQuery,
		tenantID,
		pq.Array(moldNos),
	)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	// -------------------------------------------------
	// 4. Build map of existing molds
	// -------------------------------------------------

	existingMap := make(map[string]struct{})

	for rows.Next() {
		var moldNo string

		if err := rows.Scan(&moldNo); err != nil {
			return 0, 0, err
		}

		existingMap[moldNo] = struct{}{}
	}

	// IMPORTANT:
	// rows.Next() can stop because of an iteration error,
	// so always check rows.Err() after the loop.
	if err := rows.Err(); err != nil {
		return 0, 0, err
	}

	// -------------------------------------------------
	// 5. Insert only non-existing molds
	// -------------------------------------------------

	insertQuery := `
		INSERT INTO mold (
			tenant_id,
			mold_no,
			description,
			cavities,
			target_shots,
			special_notes,
			created_by,
			updated_by
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$7
		)
	`

	for _, m := range req.Molds {

		// -------------------------------------------------
		// Skip mold if it already exists
		// -------------------------------------------------

		if _, found := existingMap[m.MoldNo]; found {
			skipped++
			continue
		}

		// -------------------------------------------------
		// Insert new mold
		// -------------------------------------------------

		_, err := tx.ExecContext(
			ctx,
			insertQuery,
			tenantID,
			m.MoldNo,
			m.Description,
			m.Cavities,
			m.TargetShots,
			m.SpecialNotes,
			userID,
		)

		if err != nil {

			// -------------------------------------------------
			// Handle duplicate insert caused by race condition
			// -------------------------------------------------

			if isUniqueViolation(err) {
				skipped++
				continue
			}

			return inserted, skipped, err
		}

		inserted++
	}

	return inserted, skipped, nil
}
