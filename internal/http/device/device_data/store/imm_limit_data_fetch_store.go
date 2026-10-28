package store

import (
	"context"
	"fmt"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Store) GetDataLimitDeviceWiseStore(
	ctx context.Context,
	req *dto.ImmDataFetchRequest,
	tenantID int64,
) (*dto.ProductionLogResponse, error) {

	rows, err := s.db.QueryContext(
		ctx,
		GetImmLimitDataDeviceWise,
		tenantID,
		req.DeviceID,
		req.MachineID,
		req.Limit,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"query IMM production records: %w",
			err,
		)
	}
	defer rows.Close()

	result := &dto.ProductionLogResponse{
		Data: make([]dto.ImmDataResponse, 0, req.Limit),
	}

	firstRow := true

	for rows.Next() {
		var item dto.ImmDataResponse

		var minCycleTimeSec float64
		var maxCycleTimeSec float64
		var avgCycleTimeSec float64

		err := rows.Scan(
			&item.ID,
			&item.EventID,
			&item.TenantID,
			&item.DeviceID,
			&item.MachineID,
			&item.MoldNo,
			&item.ProductionCount,
			&item.CycleTimeSec,
			&item.ProductionTime,
			&item.Reset,
			&item.CreatedAt,
			&item.UpdatedAt,
			&result.MinCycleTimeSec,
			&result.MaxCycleTimeSec,
			&result.AvgCycleTimeSec,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"scan IMM production record: %w",
				err,
			)
		}

		result.Data = append(result.Data, item)

		// Statistics are identical on every SQL row.
		// Store them once in the response.
		if firstRow {
			result.MinCycleTimeSec = minCycleTimeSec
			result.MaxCycleTimeSec = maxCycleTimeSec
			result.AvgCycleTimeSec = avgCycleTimeSec
			firstRow = false
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate IMM production records: %w",
			err,
		)
	}

	return result, nil
}

// func (s *Store) GetDataLimitDeviceWiseStore(ctx context.Context, req *dto.ImmDataFetchRequest, tenantID int64) ([]dto.ImmDataResponse, error) {

// 	// Data fetch from the db

// 	rows, err := s.db.QueryContext(
// 		ctx,
// 		GetImmLimitDataDeviceWise,
// 		tenantID,
// 		req.DeviceID,
// 		req.MachineID,
// 		req.Limit,
// 	)

// 	if err != nil {
// 		return nil, fmt.Errorf("query IMM prodcution records:%w", err)
// 	}
// 	defer rows.Close()

// 	result := make([]dto.ImmDataResponse, 0, req.Limit)

// 	for rows.Next() {
// 		var item dto.ImmDataResponse

// 		err := rows.Scan(
// 			&item.ID,
// 			&item.EventID,
// 			&item.TenantID,
// 			&item.DeviceID,
// 			&item.MachineID,
// 			&item.MoldNo,
// 			&item.ProductionCount,
// 			&item.CycleTimeSec,
// 			&item.ProductionTime,
// 			&item.Reset,
// 			&item.CreatedAt,
// 			&item.UpdatedAt,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("scan IMM production data: %w", err)
// 		}
// 		result = append(result, item)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("scan IMM production data: %w", err)
// 	}

// 	return result, nil
// }
