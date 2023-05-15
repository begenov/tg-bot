package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/tg-bot/internal/models"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Create(ctx context.Context, user models.User) error {
	stmt := `INSERT INTO user_profile (chat_id, name, phone, language, role, age, gender) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, stmt, user.ChatID, user.Name, user.Phone, user.Lang, user.Aim, user.Age, user.Gender)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepository) UserByChatID(ctx context.Context, chatID int) (*models.User, error) {
	stmt := `SELECT id, chat_id, name, phone, language, role, age, gender FROM user_profile WHERE chat_id = $1;`

	var user models.User
	err := r.db.QueryRowContext(ctx, stmt, chatID).Scan(&user.ID, &user.ChatID, &user.Name, &user.Phone, &user.Lang, &user.Aim, &user.Age, &user.Gender)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
