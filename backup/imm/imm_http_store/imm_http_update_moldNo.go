package immhttpstore

import (
	"context"
	"database/sql"

	immerror "github.com/rajeshbond/smart/internal/http/imm/imm_error"
)

// ============================================================
// UPDATE DEVICE MOLD
// ============================================================

func (s *ImmHttpStore) UpdateMoldNoTx(
	ctx context.Context,
	tx *sql.Tx,
	tenantID int64,
	deviceID string,
	moldNo string,
) error {

	result, err := tx.ExecContext(
		ctx,
		`
		UPDATE device_master
		SET
			mold_no = $1,
			updated_at = NOW()
		WHERE tenant_id = $2
		  AND device_id = $3
		`,
		moldNo,
		tenantID,
		deviceID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return immerror.ErrDeviceNotFound
	}

	return nil
}
