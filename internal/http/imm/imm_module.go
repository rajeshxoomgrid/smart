package imm

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	immhandler "github.com/rajeshbond/smart/internal/http/imm/imm_handler"
	immservice "github.com/rajeshbond/smart/internal/http/imm/imm_service"
	immstore "github.com/rajeshbond/smart/internal/http/imm/imm_store"
)

type ImmModule struct {
	ImmStore   *immstore.ImmStore
	ImmService *immservice.Immservice
	ImmHandler *immhandler.ImmHandler
	tokenAuth  *jwtauth.JWTAuth
}

func NewImmModule(db *sql.DB, tokenauth *jwtauth.JWTAuth) *ImmModule {
	immStore := immstore.NewImmStore(db)
	immService := immservice.NewImmStoreImmSerive(immStore)
	immHandler := immhandler.NewImmHandler(immService, tokenauth)

	return &ImmModule{
		tokenAuth:  tokenauth,
		ImmStore:   immStore,
		ImmService: immService,
		ImmHandler: immHandler,
	}
}
