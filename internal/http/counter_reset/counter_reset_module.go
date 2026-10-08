package counterreset

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	counterresethandler "github.com/rajeshbond/smart/internal/http/counter_reset/counter_reset_handler"
	counterresetservice "github.com/rajeshbond/smart/internal/http/counter_reset/counter_reset_service"
	counterresetstore "github.com/rajeshbond/smart/internal/http/counter_reset/counter_reset_store"
)

type CounterRestModule struct {
	CounterResetStore   *counterresetstore.CounterResetStore
	CounterResetService *counterresetservice.CounterRestService
	CounterResetHandler *counterresethandler.CounterResetHandler
	tokenAuth           *jwtauth.JWTAuth
}

func NewCounterResetModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *CounterRestModule {
	counterResetStore := counterresetstore.NewCounterRestStore(db)
	counterResetService := counterresetservice.NewCounterRestService(counterResetStore)
	counterresethandler := counterresethandler.NewCounterResetHandler(counterResetService, tokenAuth)

	return &CounterRestModule{
		tokenAuth:           tokenAuth,
		CounterResetStore:   counterResetStore,
		CounterResetService: counterResetService,
		CounterResetHandler: counterresethandler,
	}
}
