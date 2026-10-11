package store

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

type ProductionStore struct {
	db *sql.DB
}

func NewProductionStore(db *sql.DB) *ProductionStore {
	if db == nil {
		panic("production database connection cannot be nil")
	}

	return &ProductionStore{db: db}
}

func (s *ProductionStore) Save(
	ctx context.Context,
	req *dto.ProductionDTO,
) (*dto.ProductionResponseDTO, error) {
	// ---------------------------------------------------------
	// 1. Validate store and request
	// ---------------------------------------------------------

	if s == nil || s.db == nil {
		return nil, fmt.Errorf("production database connection is nil")
	}

	if req == nil {
		return nil, fmt.Errorf("production request is nil")
	}

	if strings.TrimSpace(req.EventID) == "" {
		return nil, fmt.Errorf("event_id is required")
	}

	if strings.TrimSpace(req.DeviceID) == "" {
		return nil, fmt.Errorf("device_id is required")
	}

	if strings.TrimSpace(req.MachineID) == "" {
		return nil, fmt.Errorf("machine_id is required")
	}

	if req.Timestamp.IsZero() {
		return nil, fmt.Errorf("production_time is required")
	}

	moldNo := "256522"
	// Enable this validation only if zero cycle time is invalid
	// for every production event in your application.
	//
	// if req.CycleTimeSec <= 0 {
	//     return nil, fmt.Errorf("cycle_time_sec must be greater than zero")
	// }

	// ---------------------------------------------------------
	// 2. Begin transaction
	// ---------------------------------------------------------

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin production transaction: %w", err)
	}
	defer tx.Rollback()

	// ---------------------------------------------------------
	// 3. Check for an existing event
	//
	// MQTT may deliver the same event more than once.
	// Return the existing row rather than inserting a duplicate.
	// ---------------------------------------------------------

	// existing, err := getProductionByEventID(
	// 	ctx,
	// 	tx,
	// 	req.EventID,
	// )
	// if err == nil {
	// 	if err := tx.Commit(); err != nil {
	// 		return nil, fmt.Errorf(
	// 			"commit existing production event: %w",
	// 			err,
	// 		)
	// 	}

	// 	return existing, nil
	// }

	// if err != sql.ErrNoRows {
	// 	return nil, fmt.Errorf(
	// 		"check existing production event: %w",
	// 		err,
	// 	)
	// }

	// ---------------------------------------------------------
	// 4. Find the mold effective at production time
	//
	// Uses the latest non-deleted mold change that was effective
	// at the time of this production event.
	// ---------------------------------------------------------

	// var moldNo sql.NullString

	// err = tx.QueryRowContext(ctx, `
	// 	SELECT new_mold_no
	// 	FROM imm_mold_change_register
	// 	WHERE tenant_id = $1
	// 	  AND device_id = $2
	// 	  AND machine_id = $3
	// 	  AND is_deleted = FALSE
	// 	  AND effective_from <= $4
	// 	ORDER BY effective_from DESC
	// 	LIMIT 1
	// `,
	// 	req.TenantID,
	// 	req.DeviceID,
	// 	req.MachineID,
	// 	req.Timestamp,
	// ).Scan(&moldNo)

	// if err != nil && err != sql.ErrNoRows {
	// 	return nil, fmt.Errorf("find effective mold: %w", err)
	// }

	// If no mold history exists, preserve an empty mold number.
	// Change this to return an error if a mold is mandatory.

	// ---------------------------------------------------------
	// 5. Reset flag
	//
	// Temporary implementation: reset history is not enabled.
	// ---------------------------------------------------------

	reset := false

	// ---------------------------------------------------------
	// 6. Insert production record
	// ---------------------------------------------------------

	_, err = tx.ExecContext(ctx, `
		INSERT INTO imm_production_log (
			event_id,
			tenant_id,
			device_id,
			machine_id,
			mold_no,
			production_count,
			cycle_time_sec,
			production_time,
			reset
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (event_id) DO NOTHING
	`,
		req.EventID,
		req.TenantID,
		req.DeviceID,
		req.MachineID,
		moldNo,
		req.Count,
		req.CycleTimeSec,
		req.Timestamp,
		reset,
	)
	if err != nil {
		return nil, fmt.Errorf("insert production record: %w", err)
	}

	// ---------------------------------------------------------
	// 7. Read back the persisted row
	//
	// Also handles a concurrent insert of the same event_id.
	// ---------------------------------------------------------

	response, err := getProductionByEventID(
		ctx,
		tx,
		req.EventID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"read saved production record: %w",
			err,
		)
	}

	// ---------------------------------------------------------
	// 8. Commit transaction
	// ---------------------------------------------------------

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf(
			"commit production record: %w",
			err,
		)
	}

	return response, nil
}

// -------------------------------------------------------------
// Row query interface
// -------------------------------------------------------------

type rowQuerier interface {
	QueryRowContext(
		context.Context,
		string,
		...any,
	) *sql.Row
}

// -------------------------------------------------------------
// Fetch a production event by event_id
// -------------------------------------------------------------

func getProductionByEventID(
	ctx context.Context,
	q rowQuerier,
	eventID string,
) (*dto.ProductionResponseDTO, error) {
	var response dto.ProductionResponseDTO

	err := q.QueryRowContext(ctx, `
		SELECT
			event_id,
			tenant_id,
			device_id,
			machine_id,
			mold_no,
			production_count,
			cycle_time_sec,
			production_time,
			reset
		FROM imm_production_log
		WHERE event_id = $1
	`,
		eventID,
	).Scan(
		&response.EventID,
		&response.TenantID,
		&response.DeviceID,
		&response.MachineID,
		&response.MoldNo,
		&response.Count,
		&response.CycleTimeSec,
		&response.Timestamp,
		&response.Reset,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// -------------------------------------------------------------
// Nullable timestamp helper
// -------------------------------------------------------------

// func nullableTime(t sql.NullTime) any {
// 	if !t.Valid {
// 		return nil

// 	}

// 	return t.Time
// }
