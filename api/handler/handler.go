package handler

import (
	"github.com/OybekAbduvosiqov/tovar/config"
	"github.com/OybekAbduvosiqov/tovar/storage"
)

type Handler struct {
	cfg     config.Config
	storage storage.StorageI
}

func NewHandler(cfg config.Config, storage storage.StorageI) *Handler {
	return &Handler{
		cfg:     cfg,
		storage: storage,
	}
}
