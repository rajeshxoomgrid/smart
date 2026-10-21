package partstore

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	partconstant "github.com/rajeshbond/smart/internal/http/part/part_constant"
	dto "github.com/rajeshbond/smart/internal/http/part/part_dto"
)

type scanner interface {
	Scan(dest ...any) error
}

func scanPart(row scanner) (*dto.PartResponse, error) {
	var p dto.PartResponse

	var color, rmGrade, description sql.NullString

	var partWt sql.NullFloat64

	var deletedBy sql.NullInt64

	err := row.Scan(
		&p.ID,
		&p.TenantID,
		&p.PartNo,
		&p.ProductName,
		&color,
		&rmGrade,
		&partWt,
		&description,
		&p.IsActive,
		&p.IsDeleted,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if color.Valid {
		p.Color = &color.String
	}
	if rmGrade.Valid {
		p.RMGrade = &rmGrade.String
	}

	if description.Valid {
		p.Description = &description.String
	}

	if partWt.Valid {
		p.PartWt = &partWt.Float64
	}

	if deletedBy.Valid {
		p.DeletedBy = &deletedBy.Int64
	}

	return &p, nil
}

func translateDBError(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" &&
			pgErr.Constraint == "uq_part_tenant_part_no_active" {
			return partconstant.ErrDuplicate
		}
	}

	return fmt.Errorf("part database operation: %w", err)
}
