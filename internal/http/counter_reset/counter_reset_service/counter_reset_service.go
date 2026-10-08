package counterresetservice

import (
	counterresetstore "github.com/rajeshbond/smart/internal/http/counter_reset/counter_reset_store"
)

type CounterRestService struct {
	CounterRestStore *counterresetstore.CounterResetStore
}

func NewCounterRestService(conterResetService *counterresetstore.CounterResetStore) *CounterRestService {
	return &CounterRestService{CounterRestStore: conterResetService}
}
