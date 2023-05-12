package services

import (
	"github.com/begenov/tg-bot/test-telegram-app/internal/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		repository: repos,
	}
}
