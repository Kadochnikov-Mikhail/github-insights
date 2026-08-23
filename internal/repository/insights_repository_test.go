package repository

import (
	"database/sql"
	"testing"

	"github-insights/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open(
		"pgx",
		"postgres://admin:password@localhost:5432/github_insights_test?sslmode=disable",
	)
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		t.Fatalf("failed to ping database: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db
}

func TestPostgresInsightsRepository_Save(t *testing.T) {
	db := setupTestDB(t)

	t.Cleanup(func() {
		_, err := db.Exec(
			"DELETE FROM github_insights WHERE username = $1",
			"test-user",
		)
		if err != nil {
			t.Errorf("failed to cleanup test data: %v", err)
		}
	})

	repo := NewInsightsRepository(db)

	insight := models.GitHubInsights{
		Username:     "test-user",
		Repositories: 5,
		TotalStars:   100,
	}

	err := repo.Save(insight)
	if err != nil {
		t.Fatalf("failed to save insight: %v", err)
	}

	var (
		username     string
		repositories int
		totalStars   int
	)

	err = db.QueryRow(`
		SELECT username, repositories, total_stars
		FROM github_insights
		WHERE username = $1
	`, insight.Username).Scan(
		&username,
		&repositories,
		&totalStars,
	)

	if err != nil {
		t.Fatalf("failed to query saved insight: %v", err)
	}

	if username != insight.Username {
		t.Errorf(
			"expected username %s, got %s",
			insight.Username,
			username,
		)
	}

	if repositories != insight.Repositories {
		t.Errorf(
			"expected repositories %d, got %d",
			insight.Repositories,
			repositories,
		)
	}

	if totalStars != insight.TotalStars {
		t.Errorf(
			"expected total stars %d, got %d",
			insight.TotalStars,
			totalStars,
		)
	}
}