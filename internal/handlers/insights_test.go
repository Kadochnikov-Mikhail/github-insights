package handlers

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github-insights/internal/models"

	"github.com/gofiber/fiber/v2"
)

type MockInsightsService struct{}

type MockInsightsServiceError struct{}

func (m MockInsightsServiceError) GetInsights(username string) (models.GitHubInsights, error) {
	return models.GitHubInsights{}, errors.New("service error")
}

func (m MockInsightsService) GetInsights(username string) (models.GitHubInsights, error) {
	return models.GitHubInsights{
		Username:     "torvalds",
		Repositories: 3,
		TotalStars:   22,
		Languages: models.RepoLangs{
			"Go":         2,
			"JavaScript": 1,
		},
	}, nil
}

func TestGetInsightsSuccess(t *testing.T) {

	app := fiber.New()

	service := MockInsightsService{}

	handler := NewInsightsHandler(service)

	app.Get("/github/insights", handler.GetInsights)

	req := httptest.NewRequest(
		"GET",
		"/github/insights?user=torvalds",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var insights models.GitHubInsights

	err = json.NewDecoder(resp.Body).Decode(&insights)

	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if insights.Username != "torvalds" {
		t.Errorf(
			"expected torvalds, got %s",
			insights.Username,
		)
	}

	if insights.Repositories != 3 {
		t.Errorf(
			"expected 3 repositories, got %d",
			insights.Repositories,
		)
	}
}

func TestGetInsightsServiceError(t *testing.T) {

	app := fiber.New()

	service := MockInsightsServiceError{}

	handler := NewInsightsHandler(service)

	app.Get("/github/insights", handler.GetInsights)

	req := httptest.NewRequest(
		"GET",
		"/github/insights?user=torvalds",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", resp.StatusCode)
	}
}

func TestHealth(t *testing.T) {

	app := fiber.New()

	app.Get("/health", Health)

	req := httptest.NewRequest(
		"GET",
		"/health",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf(
			"expected 200, got %d",
			resp.StatusCode,
		)
	}
}

