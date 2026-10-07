package immhttpstore

import (
	"context"
	"database/sql"
	"errors"

	immdto "github.com/rajeshbond/smart/internal/http/imm/imm_dto"
	immerror "github.com/rajeshbond/smart/internal/http/imm/imm_error"
)

// ============================================================
// GET DEVICE WITHOUT TENANT
// Used by MQTT ACK.
// ============================================================

func (s *ImmHttpStore) GetDeviceByIDTx(
	ctx context.Context,
	tx *sql.Tx,
	deviceID string,
) (*immdto.Device, error) {

	device := &immdto.Device{}

	err := tx.QueryRowContext(
		ctx,
		`
		SELECT
			device_id,
			machine_id,
			tenant_id,
			device_type,
			status,
			COALESCE(mold_no, '')
		FROM device_master
		WHERE device_id = $1
		FOR UPDATE
		`,
		deviceID,
	).Scan(
		&device.DeviceID,
		&device.MachineID,
		&device.TenantID,
		&device.DeviceType,
		&device.Status,
		&device.MoldNo,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, immerror.ErrDeviceNotFound
		}

		return nil, err
	}

	return device, nil
}
