package immhttphandler

import (
	"github.com/go-chi/jwtauth/v5"
	immhttpservice "github.com/rajeshbond/smart/internal/http/imm/imm_http_service"
)

type ImmHttpHandler struct {
	ImmHttpService *immhttpservice.ImmHttpService
	tokenAuth      *jwtauth.JWTAuth
}

func NewImmHttpHandler(service *immhttpservice.ImmHttpService, tokenauth *jwtauth.JWTAuth) *ImmHttpHandler {
	return &ImmHttpHandler{
		tokenAuth:      tokenauth,
		ImmHttpService: service,
	}
}
