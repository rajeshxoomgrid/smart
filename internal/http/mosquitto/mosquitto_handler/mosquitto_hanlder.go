package mosquittohandler

import (
	"github.com/go-chi/jwtauth/v5"
	mosquittoservice "github.com/rajeshbond/smart/internal/http/mosquitto/mosquitto_service"
)

type MosquittoHandler struct {
	MosquittoServide *mosquittoservice.MosquittoService
	tokenAuth        *jwtauth.JWTAuth
}

func NewMosquittoHandler(tokenauth *jwtauth.JWTAuth, moquittoService *mosquittoservice.MosquittoService) *MosquittoHandler {
	return &MosquittoHandler{
		tokenAuth:        tokenauth,
		MosquittoServide: moquittoService,
		
	}
}
