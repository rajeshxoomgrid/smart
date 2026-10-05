package immhttpservice

import immhttpstore "github.com/rajeshbond/smart/internal/http/imm/imm_http_store"

type ImmHttpService struct {
	ImmHttpService *immhttpstore.ImmHttpStore
}

func NewImmHttpService(store *immhttpstore.ImmHttpStore) *ImmHttpService {
	return &ImmHttpService{ImmHttpService: store}
}
