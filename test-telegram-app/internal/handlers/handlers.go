package handlers

import "github.com/begenov/tg-bot/test-telegram-app/internal/services"

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}
