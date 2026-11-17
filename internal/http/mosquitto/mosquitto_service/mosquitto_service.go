package mosquittoservice

import mosquittostore "github.com/rajeshbond/smart/internal/http/mosquitto/mosquitto_store"

type MosquittoService struct {
	MosquittoStore *mosquittostore.MosquittoStore
}

func NewMosquittoService(mosquittoStore *mosquittostore.MosquittoStore) *MosquittoService {
	return &MosquittoService{MosquittoStore: mosquittoStore}
}
