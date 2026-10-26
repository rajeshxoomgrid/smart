package service

import (
	"context"
	"fmt"
	"log"

	"github.com/rajeshbond/smart/internal/common/utils"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (ser *Service) GetImmLatestProductionCount(ctx context.Context, req dto.ImmDataFetchRequest, tenantID int64) (*dto.ImmDataFetechResponse, error) {

	// Latest production matcing device , station and variant

	item, err := ser.Store.GetImmTenantLatestCountByDeviceIDStore(ctx, req, tenantID)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, fmt.Errorf("production data not found")
	}

	istProductionTime := utils.ToISTNew(item.ProductionTime)

	// Shift information.
	shift, err := ser.shiftProvider.GetShiftByProductionTime(
		ctx,
		tenantID,
		istProductionTime,
	)
	if err != nil {
		log.Println("Shift Lookup Error:", err)
		return nil, err
	}

	if shift == nil {
		return nil, fmt.Errorf("shift not found")
	}

	// First counter of the current shift

	shiftFirstCount, err := ser.Store.GetImmFirstProductionInShift(ctx, req, tenantID, shift.ShiftStart, shift.ShiftEnd)

	// Calculate shift production.
	currentProductionCount := item.ProductionCount

	item.ProductionCount = currentProductionCount - shiftFirstCount
	if item.ProductionCount < 0 {
		return nil, fmt.Errorf("Issue with the data")
	}

	item.ProductionCount = item.ProductionCount + 1

	item.ShiftName = shift.ShiftName

	// Shift OEE.
	item.OEE = CalculateLiveOEE(
		shift.ShiftStart,
		item.ProductionTime,
		item.ProductionCount,
	)

	// Production day information.
	dayInfo, err := ser.shiftProvider.GetProductionDay(
		ctx,
		tenantID,
		istProductionTime,
	)
	if err != nil {
		return nil, err
	}

	// First counter for the production day.

	dayFirstCount, err := ser.Store.GetImmFirstProductionCount(ctx, tenantID, req, dayInfo.DayStart, dayInfo.DayEnd)

	if err != nil {
		return nil, err
	}

	// Calculate daily production.

	item.DayProductionCount = currentProductionCount - dayFirstCount
	if item.DayProductionCount < 0 {
		return nil, fmt.Errorf("Issue with the data")
	}

	item.DayProductionCount = item.DayProductionCount + 1

	// Daily OEE.
	item.DayOEE = CalculateLiveOEE(
		dayInfo.DayStart,
		item.ProductionTime,
		item.DayProductionCount,
	)

	// fmt.Println("===================================================")
	// fmt.Println(item.DayProductionCount)
	// fmt.Println("===================================================")

	return item, nil
}
