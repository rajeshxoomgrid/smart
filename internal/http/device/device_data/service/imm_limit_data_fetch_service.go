package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

var ErrInvalidRequest = errors.New("invalid IMM production request")

func (s *Service) GetDataLimitDeviceWise(
	ctx context.Context,
	req *dto.ImmDataFetchRequest,
	tenantID int64,
) (*dto.ProductionLogResponse, error) {

	if req == nil {
		return nil, fmt.Errorf(
			"%w: request is required",
			ErrInvalidRequest,
		)
	}

	req.DeviceID = strings.TrimSpace(req.DeviceID)
	req.MachineID = strings.TrimSpace(req.MachineID)

	if req.DeviceID == "" {
		return nil, fmt.Errorf(
			"%w: device_id is required",
			ErrInvalidRequest,
		)
	}

	if req.MachineID == "" {
		return nil, fmt.Errorf(
			"%w: machine_id is required",
			ErrInvalidRequest,
		)
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Limit < 1 || req.Limit > 100 {
		return nil, fmt.Errorf(
			"%w: limit must be between 1 and 100",
			ErrInvalidRequest,
		)
	}

	return s.Store.GetDataLimitDeviceWiseStore(
		ctx,
		req,
		tenantID,
	)
}

// func (ser *Service) GetDataLimitDeviceWiseService(ctx context.Context, req *dto.ImmDataFetchRequest, tenantID int64) (*dto.ProductionLogResponse, error) {

// 	if tenantID <= 0 {
// 		return nil, fmt.Errorf("invalid tenant ID")
// 	}

// 	req.DeviceID = strings.TrimSpace(req.DeviceID)
// 	if req.DeviceID == "" {
// 		return nil, fmt.Errorf("device ID is required")
// 	}

// 	req.MachineID = strings.TrimSpace(req.MachineID)
// 	if req.MachineID == "" {
// 		return nil, fmt.Errorf("Machine ID is required")
// 	}

// 	if req.Limit <= 0 {
// 		req.Limit = 10
// 	}

// 	if req.Limit < 1 || req.Limit > 100 {
// 		return nil, fmt.Errorf("limit must be between 1 and 100")
// 	}

// 	records, err := ser.Store.GetDataLimitDeviceWiseStore(ctx, req, tenantID)

// 	if err != nil {
// 		return nil, fmt.Errorf("get IMM production data: %w", err)
// 	}

// 	return &dto.ProductionLogResponse{
// 		Data: records,
// 	}, nil

// }
