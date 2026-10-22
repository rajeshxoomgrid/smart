package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

// ============================================================
// SAVE PREVIOUS HOUR
// ============================================================

func (s *Service) SavePreviousHour(ctx context.Context, req dto.SaveHourlyProductionRequest) (*dto.HourlyProduction, error) {

	// ========================================================
	// VALIDATION
	// ========================================================

	if s == nil || s.Store == nil {
		return nil, fmt.Errorf(
			"hourly production service/store is nil",
		)
	}

	if req.TenantID == "" {
		return nil, fmt.Errorf(
			"tenant_id is required",
		)
	}

	if req.DeviceID == "" {
		return nil, fmt.Errorf(
			"device_id is required",
		)
	}

	if req.MachineID == "" {
		return nil, fmt.Errorf(
			"machine_id is required",
		)
	}

	if req.Station == "" {
		return nil, fmt.Errorf(
			"station is required",
		)
	}

	// ========================================================
	// CUSTOMER ID
	// ========================================================

	customerID := ""

	if req.CustomerID != nil {
		customerID = *req.CustomerID
	}

	// ========================================================
	// VARIANT
	// ========================================================

	variant := ""

	if req.Variant != nil {
		variant = *req.Variant
	}

	// ========================================================
	// IST LOCATION
	// ========================================================

	istLocation, err := time.LoadLocation(
		"Asia/Kolkata",
	)

	if err != nil {
		return nil, fmt.Errorf(
			"load IST timezone: %w",
			err,
		)
	}

	// ========================================================
	// CURRENT TIME
	// ========================================================

	now := time.Now().In(istLocation)

	// ========================================================
	// BEGIN TRANSACTION
	// ========================================================

	tx, err := s.Store.BeginTx(
		ctx,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"begin transaction: %w",
			err,
		)
	}

	committed := false

	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	// ========================================================
	// CURRENT SHIFT
	// ========================================================

	shift, err := s.Store.GetCurrentShift(
		ctx,
		tx,
		req.TenantID,
		now,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"get current shift: %w",
			err,
		)
	}

	if shift == nil {
		return nil, fmt.Errorf(
			"current shift is nil",
		)
	}

	// ========================================================
	// PREVIOUS HOUR SLOT
	// ========================================================

	slot, err := s.Store.GetPreviousHourSlot(
		ctx,
		tx,
		req.TenantID,
		shift.ShiftTimingID,
		now,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"get previous hour slot: %w",
			err,
		)
	}

	if slot == nil {
		return nil, fmt.Errorf(
			"previous hour slot is nil",
		)
	}

	// ========================================================
	// PRODUCTION DAY
	// ========================================================
	//
	// Normal shift:
	//
	// 08:00 -> 20:00
	//
	// Production day = current date
	//
	// Overnight shift:
	//
	// 20:00 -> 08:00
	//
	// At 02:00:
	//
	// Production day = previous date
	//
	// ========================================================

	productionDay := now

	if isOvernightShift(shift) {

		currentClock := clockSeconds(now)

		shiftStartClock := clockSeconds(
			shift.ShiftStart,
		)

		if currentClock < shiftStartClock {

			productionDay = productionDay.AddDate(
				0,
				0,
				-1,
			)
		}
	}

	// ========================================================
	// NORMALIZE PRODUCTION DAY
	// ========================================================

	productionDay = time.Date(
		productionDay.Year(),
		productionDay.Month(),
		productionDay.Day(),
		0,
		0,
		0,
		0,
		istLocation,
	)

	// ========================================================
	// CURRENT SHIFT START
	// ========================================================
	//
	// This is required by GetProductionStats().
	//
	// Example:
	//
	// Shift A:
	// 08:00 -> 20:00
	//
	// currentShiftStart:
	// 2026-08-17 08:00:00 +05:30
	//
	// Shift B:
	// 20:00 -> 08:00
	//
	// At 02:00:
	//
	// currentShiftStart:
	// 2026-08-16 20:00:00 +05:30
	//
	// ========================================================

	currentShiftStart := combineDateAndTime(
		productionDay,
		shift.ShiftStart,
		istLocation,
	)

	// ========================================================
	// ACTUAL START
	// ========================================================

	actualStart := combineDateAndTime(
		productionDay,
		slot.SlotStart,
		istLocation,
	)

	// ========================================================
	// ACTUAL END
	// ========================================================

	actualEnd := combineDateAndTime(
		productionDay,
		slot.SlotEnd,
		istLocation,
	)

	// ========================================================
	// CROSS MIDNIGHT
	// ========================================================

	if !actualEnd.After(actualStart) {

		actualEnd = actualEnd.AddDate(
			0,
			0,
			1,
		)
	}

	// ========================================================
	// VALIDATE TIME RANGE
	// ========================================================

	if !actualEnd.After(actualStart) {

		return nil, fmt.Errorf(
			"invalid hourly slot: start=%v end=%v",
			actualStart,
			actualEnd,
		)
	}

	// ========================================================
	// GET PRODUCTION STATISTICS
	// ========================================================
	//
	// IMPORTANT:
	//
	// currentShiftStart is the NEW argument.
	//
	// ========================================================

	stats, err := s.Store.GetProductionStats(
		ctx,
		tx,

		req.TenantID,
		req.DeviceID,
		req.Station,

		currentShiftStart,
		actualStart,
		actualEnd,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"get production stats: %w",
			err,
		)
	}

	if stats == nil {
		return nil, fmt.Errorf(
			"production stats is nil",
		)
	}

	// ========================================================
	// CREATE HOURLY PRODUCTION
	// ========================================================

	item := &dto.HourlyProduction{

		// ----------------------------------------------------
		// TENANT
		// ----------------------------------------------------

		TenantID: req.TenantID,

		// ----------------------------------------------------
		// CUSTOMER
		// ----------------------------------------------------

		CustomerID: customerID,

		// ----------------------------------------------------
		// DEVICE
		// ----------------------------------------------------

		DeviceID: req.DeviceID,

		// ----------------------------------------------------
		// MACHINE
		// ----------------------------------------------------

		MachineID: req.MachineID,

		// ----------------------------------------------------
		// STATION
		// ----------------------------------------------------

		Station: req.Station,

		// ----------------------------------------------------
		// VARIANT
		// ----------------------------------------------------

		Variant: variant,

		// ----------------------------------------------------
		// PRODUCTION DAY
		// ----------------------------------------------------

		ProductionDay: productionDay,

		// ----------------------------------------------------
		// SHIFT
		// ----------------------------------------------------

		TenantShiftID: shift.TenantShiftID,

		ShiftTimingID: shift.ShiftTimingID,

		// ----------------------------------------------------
		// SLOT
		// ----------------------------------------------------

		ShiftHourSlotID: slot.ShiftHourSlotID,

		SlotIndex: slot.SlotIndex,

		SlotStart: slot.SlotStart,

		SlotEnd: slot.SlotEnd,

		// ----------------------------------------------------
		// ACTUAL TIME
		// ----------------------------------------------------

		ActualStart: actualStart,

		ActualEnd: actualEnd,

		// ----------------------------------------------------
		// PRODUCTION COUNTS
		// ----------------------------------------------------
		//
		// stats.StartProductionCount is *int
		//
		// HourlyProduction.StartProductionCount is int
		//
		// Therefore we safely dereference it.
		//
		// ----------------------------------------------------

		StartProductionCount: intValue(
			stats.StartProductionCount,
		),

		EndProductionCount: intValue(
			stats.EndProductionCount,
		),

		HourlyProductionCount: stats.HourlyProductionCount,

		// ----------------------------------------------------
		// CYCLE
		// ----------------------------------------------------

		CycleCount: stats.CycleCount,

		MinCycleTimeSec: stats.MinCycleTimeSec,

		AvgCycleTimeSec: stats.AvgCycleTimeSec,

		MaxCycleTimeSec: stats.MaxCycleTimeSec,
	}

	// ========================================================
	// DEBUG LOG
	// ========================================================

	fmt.Printf(
		"========== HOURLY PRODUCTION ==========\n"+
			"TenantID: %s\n"+
			"CustomerID: %s\n"+
			"DeviceID: %s\n"+
			"MachineID: %s\n"+
			"Station: %s\n"+
			"Variant: %s\n"+
			"ProductionDay: %v\n"+
			"CurrentShiftStart: %v\n"+
			"ShiftID: %d\n"+
			"ShiftTimingID: %d\n"+
			"SlotID: %d\n"+
			"SlotIndex: %d\n"+
			"SlotStart: %v\n"+
			"SlotEnd: %v\n"+
			"ActualStart: %v\n"+
			"ActualEnd: %v\n"+
			"StartCount: %d\n"+
			"EndCount: %d\n"+
			"HourlyCount: %d\n"+
			"CycleCount: %d\n"+
			"MinCycle: %.3f\n"+
			"AvgCycle: %.3f\n"+
			"MaxCycle: %.3f\n"+
			"=======================================\n",

		item.TenantID,
		item.CustomerID,
		item.DeviceID,
		item.MachineID,
		item.Station,
		item.Variant,

		item.ProductionDay,

		currentShiftStart,

		item.TenantShiftID,
		item.ShiftTimingID,

		item.ShiftHourSlotID,
		item.SlotIndex,

		item.SlotStart,
		item.SlotEnd,

		item.ActualStart,
		item.ActualEnd,

		item.StartProductionCount,
		item.EndProductionCount,
		item.HourlyProductionCount,

		item.CycleCount,

		item.MinCycleTimeSec,
		item.AvgCycleTimeSec,
		item.MaxCycleTimeSec,
	)

	// ========================================================
	// SAVE HOURLY PRODUCTION
	// ========================================================

	err = s.Store.SaveHourlyProduction(
		ctx,
		tx,
		item,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"save hourly production: %w",
			err,
		)
	}

	// ========================================================
	// COMMIT
	// ========================================================

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf(
			"commit hourly production: %w",
			err,
		)
	}

	committed = true

	return item, nil
}

// ============================================================
// INT POINTER HELPER
// ============================================================
//
// Converts:
//
//	*int -> int
//
// NULL becomes:
//
//	0
//
// ============================================================

func intValue(v *int) int {

	if v == nil {
		return 0
	}

	return *v
}

// ============================================================
// CHECK OVERNIGHT SHIFT
// ============================================================

func isOvernightShift(
	shift *dto.ShiftInfoHR,
) bool {

	if shift == nil {
		return false
	}

	start := clockSeconds(
		shift.ShiftStart,
	)

	end := clockSeconds(
		shift.ShiftEnd,
	)

	return start > end
}

// ============================================================
// CLOCK TO SECONDS
// ============================================================

func clockSeconds(
	t time.Time,
) int {

	return t.Hour()*3600 +
		t.Minute()*60 +
		t.Second()
}

// ============================================================
// COMBINE DATE + TIME
// ============================================================

func combineDateAndTime(
	date time.Time,
	clock time.Time,
	location *time.Location,
) time.Time {

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),

		clock.Hour(),
		clock.Minute(),
		clock.Second(),
		clock.Nanosecond(),

		location,
	)
}
