package mosquitto

import (
	"database/sql"

	"github.com/go-chi/jwtauth/v5"
	mosquittohandler "github.com/rajeshbond/smart/internal/http/mosquitto/mosquitto_handler"
	mosquittoservice "github.com/rajeshbond/smart/internal/http/mosquitto/mosquitto_service"
	mosquittostore "github.com/rajeshbond/smart/internal/http/mosquitto/mosquitto_store"
)

type MosquittoModule struct {
	MosquittoStore   *mosquittostore.MosquittoStore
	MosquittoService *mosquittoservice.MosquittoService
	MosquittoHanlder *mosquittohandler.MosquittoHandler
	tokenAuth        *jwtauth.JWTAuth
}

func NewMosquittoModule(db *sql.DB, token *jwtauth.JWTAuth) *MosquittoModule {
	mosquittoStore := mosquittostore.NewMosquittoStore(db)
	mosquittoService := mosquittoservice.NewMosquittoService(mosquittoStore)
	mosquittoHandler := mosquittohandler.NewMosquittoHandler(token, mosquittoService)

	return &MosquittoModule{
		tokenAuth:        token,
		MosquittoStore:   mosquittoStore,
		MosquittoService: mosquittoService,
		MosquittoHanlder: mosquittoHandler,
	}

}
