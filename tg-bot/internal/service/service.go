package service

import "github.com/begenov/tg-bot/internal/repository"

type Service struct {
	repository repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repository: *repos,
	}
}
