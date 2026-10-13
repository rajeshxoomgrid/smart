package part

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	parthandler "github.com/rajeshbond/smart/internal/http/part/part_handler"
	partservice "github.com/rajeshbond/smart/internal/http/part/part_service"
	partstore "github.com/rajeshbond/smart/internal/http/part/part_store"
)

type PartModule struct {
	PartStore   *partstore.PartStore
	PartService *partservice.PartService
	PartHandler *parthandler.PartHandler
	tokenAuth   *jwtauth.JWTAuth
}

func NewPartModule(db *sql.DB, tokenAuth *jwtauth.JWTAuth) *PartModule {
	partStore := partstore.NewPartStore(db)
	partService := partservice.NewPartService(partStore)
	parthandler := parthandler.NewPartHandler(partService, tokenAuth)

	return &PartModule{
		tokenAuth:   tokenAuth,
		PartStore:   partStore,
		PartService: partService,
		PartHandler: parthandler,
	}
}
