package moldchangeregister

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	moldchangehandler "github.com/rajeshbond/smart/internal/http/mold_change_register/mold_change_handler"
	moldchangeservice "github.com/rajeshbond/smart/internal/http/mold_change_register/mold_change_service"
	moldchangestore "github.com/rajeshbond/smart/internal/http/mold_change_register/mold_change_store"
)

type MoldChangeModule struct {
	MoldChangeStore   *moldchangestore.MoldChangeStore
	MoldChangeServide *moldchangeservice.MoldChangeService
	MoldChnageHandler *moldchangehandler.MoldChangeHandler
	tokenAuth         *jwtauth.JWTAuth
}

func NewMoldChangeModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *MoldChangeModule {
	moldChangeStore := moldchangestore.NewMoldChangeStore(db)
	moldChangeService := moldchangeservice.NewMoldChangeService(moldChangeStore)
	moldChangeHandler := moldchangehandler.NewMoldChangeHandler(moldChangeService, tokenAuth)

	return &MoldChangeModule{
		tokenAuth:         tokenAuth,
		MoldChangeStore:   moldChangeStore,
		MoldChangeServide: moldChangeService,
		MoldChnageHandler: moldChangeHandler,
	}
}
