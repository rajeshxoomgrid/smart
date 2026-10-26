package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Store) GetImmTenantLatestCountByDeviceIDStore(ctx context.Context, req dto.ImmDataFetchRequest, tenantID int64) (*dto.ImmDataFetechResponse, error) {
	var item dto.ImmDataFetechResponse

	err := s.db.QueryRowContext(
		ctx,
		GetLatestImmTenantCountByDeviceID,
		tenantID,
		req.DeviceID,
		req.MachineID,
		1,
	).Scan(
		&item.ID,
		&item.EventID,
		&item.TenantID,
		&item.DeviceID,
		&item.DeviceID,
		&item.MoldNo,
		&item.ProductionCount,
		&item.CycleTimeSec,
		&item.ProductionTime,
		&item.Reset,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *Store) GetImmFirstProductionInShift(
	ctx context.Context,
	req dto.ImmDataFetchRequest,
	tenantID int64,
	shiftStartTime time.Time,
	shiftEndTime time.Time,
) (int64, error) {

	var shiftFirstCount int64

	err := s.db.QueryRowContext(
		ctx,
		GetImmshiftFirstProductionCount,
		tenantID,
		req.DeviceID,
		req.MachineID,
		shiftStartTime,
		shiftEndTime,
	).Scan(&shiftFirstCount)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println("No production record found for this shift")
		return 0, nil
	}

	if err != nil {
		return 0, fmt.Errorf(
			"failed to fetch first production count: %w",
			err,
		)
	}

	// log.Println("================================")
	// log.Println("Tenant             :", tenantID)
	// log.Println("Device             :", req.DeviceID)
	// log.Println("Machine            :", req.MachineID)
	// log.Println("Shift Start        :", shiftStartTime)
	// log.Println("Shift End          :", shiftEndTime)
	// log.Println("First Production   :", shiftFirstCount)
	// log.Println("================================")

	return shiftFirstCount, nil
}

func (s *Store) GetImmFirstProductionCount(ctx context.Context, tenantID int64, req dto.ImmDataFetchRequest, dayStart time.Time, dayEnd time.Time) (int64, error) {

	var dayCount int64

	err := s.db.QueryRowContext(
		ctx,
		GetImmDayFirstProductionCount,
		tenantID,
		req.DeviceID,
		req.MachineID,
		dayStart,
		dayEnd,
	).Scan(&dayCount)

	if err != nil {

		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, err
	}

	// log.Println("================================")
	// log.Println("Tenant            :", tenantID)
	// log.Println("Device            :", req.DeviceID)
	// log.Println("Production Day    :", dayStart, "->", dayEnd)
	// log.Println("First Day Counter :", dayCount)
	// log.Println("================================")

	return dayCount, nil

}
