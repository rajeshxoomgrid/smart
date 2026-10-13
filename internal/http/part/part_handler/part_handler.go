package parthandler

import (
	"github.com/go-chi/jwtauth/v5"
	partservice "github.com/rajeshbond/smart/internal/http/part/part_service"
)

type PartHandler struct {
	PartService *partservice.PartService
	tokenAuth   *jwtauth.JWTAuth
}

func NewPartHandler(partService *partservice.PartService, tokenAuth *jwtauth.JWTAuth) *PartHandler {
	return &PartHandler{tokenAuth: tokenAuth, PartService: partService}
}
