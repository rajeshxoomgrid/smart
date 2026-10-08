package immhandler

import (
	"github.com/go-chi/jwtauth/v5"
	immservice "github.com/rajeshbond/smart/internal/http/imm/imm_service"
)

type ImmHandler struct {
	ImmService *immservice.Immservice
	tokenAuth  *jwtauth.JWTAuth
}

func NewImmHandler(ImmService *immservice.Immservice, tokenAuth *jwtauth.JWTAuth) *ImmHandler {
	return &ImmHandler{
		tokenAuth:  tokenAuth,
		ImmService: ImmService,
	}
}
