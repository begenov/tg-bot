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
	stmt := `INSERT INTO job_seeker (user_profile_id, sphere, profession, salary) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, stmt, job.ChatID, job.Sphere, job.Profession, job.Salary)
	if err != nil {
		return err
	}
	return nil
}

func (r *JobSeekerRepo) JobSeekerByChatID(ctx context.Context, chatID int) (*models.JobSeeker, error) {
	stmt := `SELECT id, user_profile_id, sphere, profession, salary FROM user_profile WHERE user_profile_id=$1;`

	var job models.JobSeeker

	if err := r.db.QueryRowContext(ctx, stmt, chatID).Scan(&job.ID, &job.ChatID, job.Sphere, job.Profession, job.Salary); err != nil {
		return nil, err
	}

	return &job, nil
}
