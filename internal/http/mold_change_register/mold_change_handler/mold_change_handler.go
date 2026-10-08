package moldchangehandler

import (
	"github.com/go-chi/jwtauth/v5"
	moldchangeservice "github.com/rajeshbond/smart/internal/http/mold_change_register/mold_change_service"
)

type MoldChangeHandler struct {
	MoldChangeService *moldchangeservice.MoldChangeService
	tokenAuth         *jwtauth.JWTAuth
}

func NewMoldChangeHandler(moldChangeServide *moldchangeservice.MoldChangeService, tokenAuth *jwtauth.JWTAuth) *MoldChangeHandler {
	return &MoldChangeHandler{
		tokenAuth:         tokenAuth,
		MoldChangeService: moldChangeServide,
	}
}
