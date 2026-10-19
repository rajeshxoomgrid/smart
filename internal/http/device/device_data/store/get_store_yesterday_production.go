package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Store) GetYesterdaysProduction1(
	ctx context.Context,
	req dto.GetProductionYesterdayRequest,
) ([]dto.ProductionShiftDTO, error) {
	station := strings.TrimSpace(req.Station)
	variant := strings.TrimSpace(req.Variant)

	// nil means the filter is optional and should not be applied.
	var stationParam any
	var variantParam any

	if station != "" {
		stationParam = station
	}

	if variant != "" {
		variantParam = variant
	}

	rows, err := s.db.QueryContext(
		ctx,
		GetYesterdayProductionQuery,
		req.Date,     // $1
		req.DeviceID, // $2
		stationParam, // $3
		variantParam, // $4
	)
	if err != nil {
		return nil, fmt.Errorf(
			"query yesterday production: %w",
			err,
		)
	}
	defer rows.Close()

	shifts := make([]dto.ProductionShiftDTO, 0, 3)

	for rows.Next() {
		var shift dto.ProductionShiftDTO

		if err := rows.Scan(
			&shift.Shift,
			&shift.ProductionCount,
		); err != nil {
			return nil, fmt.Errorf(
				"scan production shift: %w",
				err,
			)
		}

		shifts = append(shifts, shift)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate production shifts: %w",
			err,
		)
	}

	return shifts, nil
}
