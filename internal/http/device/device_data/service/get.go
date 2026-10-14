package service

import (
	"context"
	"fmt"
	"log"

	"github.com/rajeshbond/smart/internal/common/utils"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (s *Service) GetProductionLogByTenantIDAndDeviceID(ctx context.Context, req dto.GetProductionRequest) ([]dto.ProductionResponse, error) {
	return s.Store.GetByTenantAndDeviceID(ctx, req)
}

func (s *Service) GetProduction(ctx context.Context, req dto.GetProductionRequest, tenantCode int64) (*dto.ProductionResponse, error) {

	// Latest production matching device, station, and variant.
	item, err := s.Store.GetTenantCountByDevice(ctx, req)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, fmt.Errorf("production data not found")
	}

	istProductionTime := utils.ToISTNew(item.ProductionTime)

	// Shift information.
	shift, err := s.shiftProvider.GetShiftByProductionTime(
		ctx,
		tenantCode,
		istProductionTime,
	)
	if err != nil {
		log.Println("Shift Lookup Error:", err)
		return nil, err
	}

	if shift == nil {
		return nil, fmt.Errorf("shift not found")
	}

	// First counter for the current shift.
	firstShiftCount, err := s.Store.GetFirstShiftProductionCount(
		ctx,
		req.TenantID,
		req.DeviceID,
		req.Station,
		req.Variant,
		shift.ShiftStart,
		shift.ShiftEnd,
	)
	if err != nil {
		return nil, err
	}

	// Calculate shift production.
	currentProductionCount := item.ProductionCount

	item.ProductionCount = currentProductionCount - firstShiftCount
	if item.ProductionCount < 0 {
		item.ProductionCount = 0
	}

	item.ShiftName = shift.ShiftName

	// Shift OEE.
	item.OEE = CalculateLiveOEE(
		shift.ShiftStart,
		item.ProductionTime,
		item.ProductionCount,
	)

	// Production day information.
	dayInfo, err := s.shiftProvider.GetProductionDay(
		ctx,
		tenantCode,
		istProductionTime,
	)
	if err != nil {
		return nil, err
	}

	// First counter for the production day.
	dayFirstCount, err := s.Store.GetFirstDayProductionCount(
		ctx,
		req.TenantID,
		req.DeviceID,
		req.Station,
		req.Variant,
		dayInfo.DayStart,
		dayInfo.DayEnd,
	)
	if err != nil {
		return nil, err
	}

	// Calculate daily production.
	item.DayProduction = currentProductionCount - dayFirstCount
	if item.DayProduction < 0 {
		item.DayProduction = 0
	}

	// Daily OEE.
	item.DayOEE = CalculateLiveOEE(
		dayInfo.DayStart,
		item.ProductionTime,
		item.DayProduction,
	)

	return item, nil
}

func (s *Service) GetProduction1(ctx context.Context, req dto.GetProductionRequest, tenantCode int64) (*dto.ProductionResponse, error) {

	//-------------------------------------------------
	// Latest Production
	//-------------------------------------------------

	item, err := s.Store.GetTenantCountByDevice(ctx, req)
	if err != nil {
		return nil, err
	}
	istNewProductionTime := utils.ToISTNew(item.ProductionTime)
	// log.Println("IST Production -----> ", istNewProductionTime)
	// log.Println("Latest Creation ----->", item.CreatedAt)
	// log.Println("Production Time ----->", item.ProductionTime)
	// log.Println("================================")

	//-------------------------------------------------
	// Shift Info
	//-------------------------------------------------

	shift, err := s.shiftProvider.GetShiftByProductionTime(
		ctx,
		tenantCode,
		// item.ProductionTime, // Use production time instead of created time
		istNewProductionTime,
	)

	if err != nil {
		log.Println("Shift Lookup Error :", err)
		return nil, err
	}

	if shift == nil {
		return nil, fmt.Errorf("shift not found")
	}

	// log.Println("========== SHIFT ==========")
	// log.Println("Shift Name :", shift.ShiftName)
	// log.Println("Shift Start:", shift.ShiftStart)
	// log.Println("Shift End  :", shift.ShiftEnd)

	// log.Println("Local Time :", time.Now())
	// log.Println("UTC Time   :", time.Now().UTC())

	//-------------------------------------------------
	// First Shift Counter
	//-------------------------------------------------

	firstShiftCount, err := s.Store.GetFirstShiftProductionCount(
		ctx,
		req.TenantID,
		req.DeviceID,
		req.Station,
		req.Variant,
		shift.ShiftStart,
		shift.ShiftEnd,
	)

	if err != nil {
		return nil, err
	}

	// log.Println("First Shift Count :", firstShiftCount)

	//-------------------------------------------------
	// Shift Production
	//-------------------------------------------------

	currentProductionCount := item.ProductionCount

	item.ProductionCount = currentProductionCount - firstShiftCount

	if item.ProductionCount < 0 {
		item.ProductionCount = 0
	}
	item.ShiftName = shift.ShiftName
	//-------------------------------------------------
	// Shift OEE
	//-------------------------------------------------

	item.OEE = CalculateLiveOEE(
		shift.ShiftStart,
		item.ProductionTime,
		item.ProductionCount,
	)

	//-------------------------------------------------
	// Production Day
	//-------------------------------------------------

	dayInfo, err := s.shiftProvider.GetProductionDay(
		ctx,
		tenantCode,
		istNewProductionTime,
	)

	if err != nil {
		return nil, err
	}

	dayFirstCount, err := s.Store.GetFirstShiftProductionCount(
		ctx,
		req.TenantID,
		req.DeviceID,
		req.Station,
		req.Variant,
		dayInfo.DayStart,
		dayInfo.DayEnd,
	)

	if err != nil {
		return nil, err
	}

	// log.Println("Day First Count ---->:", dayFirstCount)

	item.DayProduction = currentProductionCount - dayFirstCount

	if item.DayProduction < 0 {
		item.DayProduction = 0
	}

	item.DayOEE = CalculateLiveOEE(
		dayInfo.DayStart,
		item.ProductionTime,
		item.DayProduction,
	)

	// //-------------------------------------------------
	// // Debug Logs
	// //-------------------------------------------------

	// log.Println("========== SHIFT ==========")
	// log.Println("Shift Name      :", shift.ShiftName)
	// log.Println("Shift Start     :", shift.ShiftStart)
	// log.Println("Shift End       :", shift.ShiftEnd)

	// log.Println("========== DAY ==========")
	// log.Println("Day Start       :", dayInfo.DayStart)
	// log.Println("Day End         :", dayInfo.DayEnd)

	// log.Println("Production Time :", item.ProductionTime)
	// log.Println("Created Time    :", item.CreatedAt)

	// log.Println("Shift Count     :", item.ProductionCount)
	// log.Println("Day Count       :", item.DayProduction)

	return item, nil
}

//
// func (s *Service) GetProduction(
// 	ctx context.Context,
// 	req dto.GetProductionRequest,
// 	tenantID int64,
// ) (*dto.ProductionResponse, error) {

// 	//-------------------------------------------------
// 	// Latest Production
// 	//-------------------------------------------------

// 	item, err := s.Store.GetTenantCountByDevice(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("Latest Creation at ----->", item.CreatedAt)
// 	log.Println("Last Counter ----->", item.ProductionTime)
// 	log.Println("==============================")
// 	//-------------------------------------------------
// 	// Shift Info
// 	//-------------------------------------------------

// 	shift, err := s.shiftProvider.GetShiftByProductionTime(
// 		ctx,
// 		tenantID,
// 		item.CreatedAt,
// 	)
// 	log.Println("Shift data -----> ", shift)
// 	// istTime := utils.ToIST(item.ProductionTime)

// 	// item.ProductionTime = istTime

// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// log.Println("Shift Name ----->", shift.ShiftName)
// 	// item.ShiftName = shift.ShiftName
// 	log.Println("Local Time", time.Now())
// 	log.Println("================================")
// 	log.Println("UTC Time", time.Now().UTC())
// 	//-------------------------------------------------
// 	// First Shift Count
// 	//-------------------------------------------------

// 	firstShiftCount, err := s.Store.GetFirstShiftProductionCount(
// 		ctx,
// 		req.TenantID,
// 		req.DeviceID,
// 		req.Station,
// 		shift.ShiftStart,
// 		shift.ShiftEnd,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Println("----> shift 1st", firstShiftCount)
// 	//-------------------------------------------------
// 	// Shift Production
// 	//-------------------------------------------------
// 	currentProductionCount := item.ProductionCount
// 	item.ProductionCount = item.ProductionCount - firstShiftCount

// 	//-------------------------------------------------
// 	// Live OEE
// 	//-------------------------------------------------

// 	log.Printf("Shift Start : %v", shift.ShiftStart)
// 	log.Printf("Shift Count : %d", item.ProductionCount)

// 	item.OEE = CalculateLiveOEE(
// 		shift.ShiftStart,
// 		item.ProductionTime,
// 		item.ProductionCount,
// 	)

// 	// Day Production

// 	dayFirst, err := s.shiftProvider.GetProductionDay(ctx, tenantID, item.ProductionTime)

// 	if err != nil {
// 		return nil, err
// 	}

// 	dayFirstCount, err := s.Store.GetFirstShiftProductionCount(ctx, req.TenantID, req.DeviceID, req.Station, dayFirst.DayStart, dayFirst.DayEnd)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dayProductionCount := currentProductionCount - dayFirstCount

// 	log.Println("====> Day Production Count ", dayProductionCount)

// 	item.DayProduction = dayProductionCount

// 	item.DayOEE = CalculateLiveOEE(
// 		dayFirst.DayStart,
// 		item.ProductionTime,
// 		item.DayProduction,
// 	)

// 	log.Println("========== SHIFT ==========")
// 	log.Println("Shift Name      :", shift.ShiftName)
// 	log.Println("Shift Start     :", shift.ShiftStart)
// 	log.Println("Shift End       :", shift.ShiftEnd)

// 	log.Println("========== DAY ==========")
// 	log.Println("Day Start       :", dayFirst.DayStart)
// 	log.Println("Day End         :", dayFirst.DayEnd)

// 	log.Println("Production Time :", item.ProductionTime)
// 	log.Println("Created Time :", item.CreatedAt)

// 	log.Println("Shift Count     :", item.ProductionCount)
// 	log.Println("Day Count       :", item.DayProduction)

// 	// log.Printf("OEE : %+v", item.OEE)

// 	return item, nil
// }

// CalculateDayProduction returns today's production count
// based on the latest machine counter.
// func (s *Service) CalculateDayProduction(
// 	ctx context.Context,
// 	req dto.GetProductionRequest,
// 	tenantID int64,
// 	latest *dto.ProductionResponse,
// ) (int64, error) {

// 	//---------------------------------------------------------
// 	// Production Day Information
// 	//---------------------------------------------------------

// 	dayInfo, err := s.shiftProvider.GetProductionDay(
// 		ctx,
// 		tenantID,
// 		latest.ProductionTime,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	//---------------------------------------------------------
// 	// First Counter of Production Day
// 	//---------------------------------------------------------

// 	firstDayCount, err := s.Store.GetFirstDayProductionCount(
// 		ctx,
// 		req.TenantID,
// 		req.DeviceID,
// 		req.Station,
// 		dayInfo.DayStart,
// 		dayInfo.DayEnd,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	//---------------------------------------------------------
// 	// Calculate Day Production
// 	//---------------------------------------------------------

// 	dayProduction := latest.ProductionCount - firstDayCount

// 	if dayProduction < 0 {
// 		dayProduction = 0
// 	}

// 	return dayProduction, nil
// }

// func (s *Service) GetProduction(ctx context.Context, req dto.GetProductionRequest, tenantID int64) (*dto.ProductionResponse, error) {
// 	//-------------------------------------------------
// 	// Latest Production
// 	//-------------------------------------------------
// 	item, err := s.Store.GetTenantCountByDevice(ctx, req)

// 	if err != nil {
// 		return nil, err
// 	}
// 	//-------------------------------------------------
// 	// Shift Info
// 	//-------------------------------------------------
// 	shift, err := s.shiftProvider.GetShiftByProductionTime(
// 		ctx,
// 		tenantID,
// 		item.ProductionTime,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	item.ShiftName = shift.ShiftName

// 	//-------------------------------------------------
// 	// First Shift Count
// 	//-------------------------------------------------

// 	firstshiftCount, err := s.Store.GetFirstShiftProductionCount(ctx, req.TenantID, req.DeviceID, req.Station, shift.ShiftStart, shift.ShiftEnd)
// 	if err != nil {
// 		return nil, err
// 	}

// 	item.ProductionCount = item.ProductionCount - firstshiftCount

// 	return item, nil

// }
