package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/tg-bot/internal/models"
)

type JobSeekerRepo struct {
	db *sql.DB
}

func NewJobSeekerRepo(db *sql.DB) *JobSeekerRepo {
	return &JobSeekerRepo{db: db}
}

func (r *JobSeekerRepo) CreateJobSeeker(ctx context.Context, job models.JobSeeker) error {
	return nil
}

func (r *JobSeekerRepo) JobSeekerByChatID(ctx context.Context, chatID int) (*models.JobSeeker, error) {
	return nil, nil
}
