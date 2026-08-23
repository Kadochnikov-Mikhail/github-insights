package repository

import (
	"database/sql"

	"github-insights/internal/models"
)

type InsightsRepository interface {
	Save(models.GitHubInsights) error
	GetByUsername(username string) ([]models.GitHubInsights, error)
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

func (r *PostgresInsightsRepository) GetByUsername(username string) ([]models.GitHubInsights, error) {

	query := `
		SELECT username, repositories, total_stars, created_at
		FROM github_insights
		WHERE username = $1
		ORDER BY created_at ASC
	`

	rows, err := r.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var insights []models.GitHubInsights

	for rows.Next() {
		var insight models.GitHubInsights

		err := rows.Scan(
			&insight.Username,
			&insight.Repositories,
			&insight.TotalStars,
			&insight.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		insights = append(insights, insight)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return insights, nil
}
