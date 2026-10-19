package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_data/helper"
)

func (s *Service) GetYesterdatsProduction1(
	ctx context.Context,
	req dto.GetProductionYesterdayRequest,
	tenantID string,
) (*dto.GetYesterdaysProductionResponse, error) {
	// Normalize request values.
	req.DeviceID = strings.TrimSpace(req.DeviceID)
	req.Station = strings.TrimSpace(req.Station)
	req.Variant = strings.TrimSpace(req.Variant)
	req.Date = strings.TrimSpace(req.Date)
	tenantID = strings.TrimSpace(tenantID)

	// Validate authenticated tenant.
	if tenantID == "" {
		return nil, fmt.Errorf(
			"%w: tenant_id is required",
			helper.ErrInvalidRequest,
		)
	}

	// Device ID is mandatory.
	if req.DeviceID == "" {
		return nil, fmt.Errorf(
			"%w: device_id is required",
			helper.ErrInvalidRequest,
		)
	}

	// Validate date format.
	productionDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: date must use YYYY-MM-DD format",
			helper.ErrInvalidRequest,
		)
	}

	req.Date = productionDate.Format("2006-01-02")

	// Retrieve shift production counts.
	shifts, err := s.Store.GetYesterdaysProduction1(
		ctx,
		req,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get production shifts: %w",
			err,
		)
	}

	// Complete Day already includes Shift A and Shift B.
	// Therefore, use Complete Day as the total rather than
	// adding all three shift counts together.
	var total int64

	for _, shift := range shifts {
		if shift.Shift == "Complete Day" {
			total = shift.ProductionCount
			break
		}
	}

	// Build API response.
	return &dto.GetYesterdaysProductionResponse{
		TenantID: tenantID,
		DeviceID: req.DeviceID,
		Station:  req.Station,
		Variant:  req.Variant,
		Date:     req.Date,
		Shifts:   shifts,
		Total:    total,
	}, nil
}
