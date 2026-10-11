package service

import (
	"context"
	"fmt"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

func (s *ImmService) Save(
	ctx context.Context,
	req *dto.ProductionDTO,
) (*dto.ProductionResponseDTO, error) {
	if s == nil || s.ImmProductionStore == nil {
		return nil, fmt.Errorf("production service is not initialized")
	}
	if req == nil {
		return nil, fmt.Errorf("production request is nil")
	}
	if req.EventID == "" {
		return nil, fmt.Errorf("event_id is required")
	}
	if req.TenantID <= 0 {
		return nil, fmt.Errorf("tenant_id must be positive")
	}
	if req.DeviceID == "" {
		return nil, fmt.Errorf("device_id is required")
	}
	if req.Count < 0 {
		return nil, fmt.Errorf("production_count cannot be negative")
	}
	if req.CycleTimeSec < 0 {
		return nil, fmt.Errorf("cycle_time_sec cannot be negative")
	}
	if req.Timestamp.IsZero() {
		return nil, fmt.Errorf("production_time is required")
	}
	// fmt.Println("<-----------------------------------Service-------------------------------------------->")
	return s.ImmProductionStore.Save(ctx, req)

	// return nil, nil

}

// ============================================================
// SAVE PRODUCTION
// ============================================================

// func (s *ImmService) Save(
// 	ctx context.Context,
// 	req *dto.ProductionDTO,
// ) error {

// 	if req == nil {
// 		return fmt.Errorf("production request is nil")
// 	}

// 	fmt.Println("<----------------------Service-------------------------->", req)
// 	// return nil
// 	return s.Save(ctx, req)
// }
