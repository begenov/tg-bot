package handler

import (
	"github.com/begenov/tg-bot/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {

	return &Handler{
		service: service,
	}
}
