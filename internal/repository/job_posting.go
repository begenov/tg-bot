package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/tg-bot/internal/models"
)

type JobPostingStorage struct {
	db *sql.DB
}

func NewJobPostingStorage(db *sql.DB) *JobPostingStorage {
	return &JobPostingStorage{
		db: db,
	}
}
func (r *JobPostingStorage) CreateJobPosting(ctx context.Context, job models.Vacancy) error {
	stmt := `INSERT INTO vacancies (user_profile_id, company, bin, sphere, position, salary, requirements, responsibilities) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, stmt, job.ChatID, job.Company, job.BIN, job.Sphere, job.Position, job.Salary, job.Requirements, job.Responsibilities)
	if err != nil {
		return err
	}
	return nil
}

func (r *JobPostingStorage) JobSeekerByChatID(ctx context.Context, chatID int) (*models.JobSeeker, error) {
	stmt := `SELECT id, user_profile_id, sphere, profession, salary FROM user_profile WHERE user_profile_id=$1;`

	var job models.JobSeeker

	if err := r.db.QueryRowContext(ctx, stmt, chatID).Scan(&job.ID, &job.ChatID, job.Sphere, job.Profession, job.Salary); err != nil {
		return nil, err
	}

	return &job, nil
}
