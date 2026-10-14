package store

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Store) GetByTenantAndDeviceID(ctx context.Context, req dto.GetProductionRequest) ([]dto.ProductionResponse, error) {
	rows, err := s.db.QueryContext(ctx, GetLatestProductionByDevice, req.TenantID, req.DeviceID, req.Station, req.Limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]dto.ProductionResponse, 0)

	for rows.Next() {
		var item dto.ProductionResponse

		err := rows.Scan(
			&item.TenantID,
			&item.EventID,
			&item.CustomerID,
			&item.DeviceID,
			&item.MachineID,
			&item.Station,
			&item.ProductionCount,
			&item.CycleTimeSec,
			&item.ProductionTime,
			&item.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *Store) GetTenantCountByDevice(ctx context.Context, req dto.GetProductionRequest) (*dto.ProductionResponse, error) {
	var item dto.ProductionResponse

	err := s.db.QueryRowContext(ctx, GetLatestProductionByDevice, req.TenantID, req.DeviceID, req.Station, req.Variant, 1).Scan(
		&item.TenantID,
		&item.EventID,
		&item.CustomerID,
		&item.DeviceID,
		&item.MachineID,
		&item.Station,
		&item.Variant,
		&item.ProductionCount,
		&item.CycleTimeSec,
		&item.ProductionTime,
		&item.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *Store) GetFirstShiftProductionCount(ctx context.Context, tenantCode string, deviceID string, station string, variant string,
	shiftStart time.Time, shiftEnd time.Time) (int64, error) {

	var count int64

	// log.Println("================================")
	// log.Println("shiftStrt time", shiftStart)
	// log.Println("shiftEnd time", shiftEnd)
	// log.Println("================================")

	err := s.db.QueryRowContext(
		ctx,
		GetFirstShiftProductionCount,
		tenantCode,
		deviceID,
		station,
		variant,
		shiftStart,
		shiftEnd,
	).Scan(&count)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}
	// log.Println("Count----->", count)
	log.Println("================================")
	log.Println("Tenant            :", tenantCode)
	log.Println("Device            :", deviceID)
	log.Println("Station           :", station)
	log.Println("Variant           :", variant)
	log.Println("Production Day    :", shiftStart, "->", shiftEnd)
	log.Println("Shifts 1st count Counter :", count)
	log.Println("================================")
	return count, nil
}

func (s *Store) GetFirstDayProductionCount(ctx context.Context, tenantCode string, deviceID string, station string, variant string, dayStart time.Time, dayEnd time.Time) (int64, error) {

	var count int64

	err := s.db.QueryRowContext(
		ctx,
		GetFirstDayProductionCount,
		tenantCode,
		deviceID,
		station,
		variant,
		dayStart,
		dayEnd,
	).Scan(&count)

	if err != nil {

		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, err
	}

	log.Println("================================")
	log.Println("Tenant            :", tenantCode)
	log.Println("Device            :", deviceID)
	log.Println("Station           :", station)
	log.Println("Variant           :", variant)
	log.Println("Production Day    :", dayStart, "->", dayEnd)
	log.Println("First Day Counter :", count)
	log.Println("================================")

	return count, nil
}

// func (s *Store) GetByTenantAndDeviceID(ctx context.Context, tenantID, deviceID string) (*dto.ProductionResponse, error) {

// 	var resp dto.ProductionResponse

// 	err := s.db.QueryRowContext(ctx, GetLatestProductionByDevice, tenantID, deviceID).Scan(
// 		&resp.TenantID,
// 		&resp.CustomerID,
// 		&resp.DeviceID,
// 		&resp.MachineID,
// 		&resp.Station,
// 		&resp.ProductionCount,
// 		&resp.CycleTimeSec,
// 		&resp.ProductionTime,
// 		&resp.CreatedAt,
// 	)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, sql.ErrNoRows
// 		}
// 		return nil, err
// 	}

// 	return &resp, nil

// }
