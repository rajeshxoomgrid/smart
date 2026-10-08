package counterresethandler

import (
	"github.com/go-chi/jwtauth/v5"
	counterresetservice "github.com/rajeshbond/smart/internal/http/counter_reset/counter_reset_service"
)

type CounterResetHandler struct {
	CounterResetService *counterresetservice.CounterRestService
	tokenAuth           *jwtauth.JWTAuth
}

func NewCounterResetHandler(counterresetservice *counterresetservice.CounterRestService, tokenAuth *jwtauth.JWTAuth) *CounterResetHandler {
	return &CounterResetHandler{
		tokenAuth:           tokenAuth,
		CounterResetService: counterresetservice,
	}
}
