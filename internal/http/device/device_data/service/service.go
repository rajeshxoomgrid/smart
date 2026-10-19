package service

import (
	"context"

	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_data/shiftprovider"
	"github.com/rajeshbond/smart/internal/http/device/device_data/store"
)

type Service struct {
	Store         *store.Store
	shiftProvider shiftprovider.ShiftProvider
}

func (s *Service) GetYesterdayProdcution(ctx context.Context, req dto.GetProductionYesterdayRequest, d int64) (any, error) {
	panic("unimplemented")
}

func NewService(store *store.Store, shiftProvider shiftprovider.ShiftProvider) *Service {
	return &Service{
		Store:         store,
		shiftProvider: shiftProvider,
	}
}
