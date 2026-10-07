package immhttp

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	immhttphandler "github.com/rajeshbond/smart/internal/http/imm/imm_http_handler"
	immhttpservice "github.com/rajeshbond/smart/internal/http/imm/imm_http_service"
	immhttpstore "github.com/rajeshbond/smart/internal/http/imm/imm_http_store"
)

type ImmHttpModule struct {
	ImmHttpStore   *immhttpstore.ImmHttpStore
	ImmHttpService *immhttpservice.ImmHttpService
	ImmHttpHandler *immhttphandler.ImmHttpHandler
	tokenAuth      *jwtauth.JWTAuth
}

func NewImmHttpModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *ImmHttpModule {
	immHttpStore := immhttpstore.NewImmHttpStore(db)
	immhttpservice := immhttpservice.NewImmHttpService(immHttpStore)
	immHttpHandler := immhttphandler.NewImmHttpHandler(immhttpservice, tokenAuth)

	return &ImmHttpModule{
		tokenAuth:      tokenAuth,
		ImmHttpStore:   immHttpStore,
		ImmHttpService: immhttpservice,
		ImmHttpHandler: immHttpHandler,
	}
}
