package repository

import (
	"database/sql"

	"github-insights/internal/models"
)

type InsightsRepository interface {
	Save(models.GitHubInsights) error
}

type PostgresInsightsRepository struct {
	db *sql.DB
}

func NewInsightsRepository(db *sql.DB) *PostgresInsightsRepository {
	return &PostgresInsightsRepository{
		db: db,
	}
}

func (r *PostgresInsightsRepository) Save(insight models.GitHubInsights) error {

	query := `
		INSERT INTO github_insights
		(username, repositories, total_stars)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		query,
		insight.Username,
		insight.Repositories,
		insight.TotalStars,
	)

	return err
}