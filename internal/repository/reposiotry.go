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

type JobSeeker interface {
	JobSeekerByChatID(ctx context.Context, chatID int) (*models.JobSeeker, error)
	CreateJobSeeker(ctx context.Context, job models.JobSeeker) error
}

type JobPosting interface {
	CreateJobPosting(ctx context.Context, job models.Vacancy) error
}

type Repository struct {
	User       UserIR
	JobSeeker  JobSeeker
	JobPosting JobPosting
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:       NewProfileRepository(db),
		JobSeeker:  NewJobSeekerRepo(db),
		JobPosting: NewJobPostingStorage(db),
	}
}
