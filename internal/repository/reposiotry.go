package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/tg-bot/internal/models"
)

type UserIR interface {
	Create(ctx context.Context, user models.User) error
	UserByChatID(ctx context.Context, chatID int) (*models.User, error)
}

type Repository struct {
	User ProfileRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		*NewProfileRepository(db),
	}
}
