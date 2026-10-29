package partservice

import (
	partstore "github.com/rajeshbond/smart/internal/http/part/part_store"
)

type PartService struct {
	PartStore *partstore.PartStore
}

func NewPartService(partStore *partstore.PartStore) *PartService {
	return &PartService{PartStore: partStore}
}
