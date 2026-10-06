package immhttpstore

import (
	"context"
	"database/sql"
	"errors"

	immdto "github.com/rajeshbond/smart/internal/http/imm/imm_dto"
	immerror "github.com/rajeshbond/smart/internal/http/imm/imm_error"
)

// ============================================================
// CREATE PENDING COMMAND
// ============================================================
// ============================================================
// CREATE PENDING COMMAND
// ============================================================

func (s *ImmHttpStore) CreatePendingTx(
	ctx context.Context,
	tx *sql.Tx,
	tenantID int64,
	userID int64,
	deviceID string,
	machineID string,
	oldMoldNo string,
	requestedMoldNo string,
) (int64, error) {

	var id int64

	err := tx.QueryRowContext(
		ctx,
		`
		INSERT INTO imm_mold_change_command (
			tenant_id,
			user_id,
			device_id,
			machine_id,
			old_mold_no,
			requested_mold_no,
			status
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			'PENDING'
		)
		RETURNING id
		`,
		tenantID,
		userID,
		deviceID,
		machineID,
		oldMoldNo,
		requestedMoldNo,
	).Scan(&id)

	return id, err
}

// ============================================================
// GET PENDING COMMAND
// LOCK ROW
// ============================================================

// ============================================================
// GET PENDING COMMAND
// ============================================================

func (s *ImmHttpStore) GetPendingForDeviceTx(
	ctx context.Context,
	tx *sql.Tx,
	tenantID int64,
	deviceID string,
) (*immdto.MoldChangeCommand, error) {

	command := &immdto.MoldChangeCommand{}

	err := tx.QueryRowContext(
		ctx,
		`
		SELECT
			id,
			tenant_id,
			user_id,
			device_id,
			machine_id,
			COALESCE(old_mold_no, ''),
			requested_mold_no,
			status,
			COALESCE(failure_reason, ''),
			created_at
		FROM imm_mold_change_command
		WHERE tenant_id = $1
		  AND device_id = $2
		  AND status = 'PENDING'
		ORDER BY id DESC
		LIMIT 1
		FOR UPDATE
		`,
		tenantID,
		deviceID,
	).Scan(
		&command.ID,
		&command.TenantID,
		&command.UserID,
		&command.DeviceID,
		&command.MachineID,
		&command.OldMoldNo,
		&command.RequestedMoldNo,
		&command.Status,
		&command.FailureReason,
		&command.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return command, nil
}

// ============================================================
// GET LATEST COMMAND
// ============================================================

func (s *ImmHttpStore) GetLatestForDevice(
	ctx context.Context,
	tenantID int64,
	deviceID string,
) (*immdto.MoldChangeCommand, error) {

	command := &immdto.MoldChangeCommand{}

	err := s.db.QueryRowContext(
		ctx,
		`
		SELECT
			id,
			tenant_id,
			user_id,
			device_id,
			machine_id,
			COALESCE(old_mold_no, ''),
			requested_mold_no,
			status,
			COALESCE(failure_reason, ''),
			created_at
		FROM imm_mold_change_command
		WHERE tenant_id = $1
		  AND device_id = $2
		ORDER BY id DESC
		LIMIT 1
		`,
		tenantID,
		deviceID,
	).Scan(
		&command.ID,
		&command.TenantID,
		&command.UserID,
		&command.DeviceID,
		&command.MachineID,
		&command.OldMoldNo,
		&command.RequestedMoldNo,
		&command.Status,
		&command.FailureReason,
		&command.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return command, nil
}

// ============================================================
// SUCCESS
// ============================================================

func (s *ImmHttpStore) MarkSuccessTx(
	ctx context.Context,
	tx *sql.Tx,
	id int64,
) error {

	result, err := tx.ExecContext(
		ctx,
		`
		UPDATE imm_mold_change_command
		SET
			status = 'SUCCESS',
			failure_reason = NULL,
			confirmed_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
		  AND status = 'PENDING'
		`,
		id,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return immerror.ErrInvalidACK
	}

	return nil
}

// ============================================================
// FAILED
// ============================================================

func (s *ImmHttpStore) MarkFailedTx(
	ctx context.Context,
	tx *sql.Tx,
	id int64,
	reason string,
) error {

	result, err := tx.ExecContext(
		ctx,
		`
		UPDATE imm_mold_change_command
		SET
			status = 'FAILED',
			failure_reason = $2,
			updated_at = NOW()
		WHERE id = $1
		  AND status = 'PENDING'
		`,
		id,
		reason,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return immerror.ErrCommandAlreadyCompleted
	}

	return nil
}

// ============================================================
// MARK OLD PENDING COMMANDS FAILED
//
// Returns number of commands changed.
//
// This is used by the 30-second timeout worker.
// ============================================================

// ============================================================
// TIMEOUT PENDING COMMANDS
// ============================================================

func (s *ImmHttpStore) MarkExpiredPendingFailed(
	ctx context.Context,
	tx *sql.Tx,
) (int64, error) {

	result, err := tx.ExecContext(
		ctx,
		`
		UPDATE imm_mold_change_command
		SET
			status = 'FAILED',
			failure_reason =
				'Machine did not confirm within 30 seconds',
			updated_at = NOW()
		WHERE status = 'PENDING'
		  AND created_at <= NOW() - INTERVAL '30 seconds'
		`,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
