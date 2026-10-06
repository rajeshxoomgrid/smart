package immhttpservice

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	immhttpstore "github.com/rajeshbond/smart/internal/http/imm/imm_http_store"
)

type ImmHttpService struct {
	ImmHttpService *immhttpstore.ImmHttpStore
	MQTT           paho.Client
}

func NewImmHttpService(store *immhttpstore.ImmHttpStore) *ImmHttpService {
	return &ImmHttpService{ImmHttpService: store}
}
