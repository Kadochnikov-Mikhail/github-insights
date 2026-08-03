package services

import (
	"errors"
	"testing"

	"github-insights/internal/apperror"
	"github-insights/internal/models"
)

func TestGetInsightsSuccess(t *testing.T) {

	client := MockGitHubClient{
		Repos: []models.GitHubRepo{
			{
				Name:     "repo1",
				Stars:    10,
				Language: "Go",
			},
			{
				Name:     "repo2",
				Stars:    5,
				Language: "Go",
			},
			{
				Name:     "repo3",
				Stars:    7,
				Language: "JavaScript",
			},
		},
	}

	service := NewInsightsService(client)

	insights, err := service.GetInsights("torvalds")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if insights.Username != "torvalds" {
		t.Errorf("expected torvalds, got %s", insights.Username)
	}

	if insights.Repositories != 3 {
		t.Errorf("expected 3, got %d", insights.Repositories)
	}

	if insights.TotalStars != 22 {
		t.Errorf("expected 22, got %d", insights.TotalStars)
	}

	if insights.Languages["Go"] != 2 {
		t.Errorf(
			"expected Go count 2, got %d",
			insights.Languages["Go"],
		)
	}

	if insights.Languages["JavaScript"] != 1 {
		t.Errorf(
			"expected JavaScript count 1, got %d",
			insights.Languages["JavaScript"],
		)
	}

}
func TestGetInsightsClientError(t *testing.T) {
	client := MockGitHubClient{
		Err: errors.New("github api error"),
	}

	service := NewInsightsService(client)

	_, err := service.GetInsights("torvalds")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetInsightsEmptyUsername(t *testing.T) {

	client := MockGitHubClient{}

	service := NewInsightsService(client)

	_, err := service.GetInsights("")

	if err != apperror.ErrUsernameRequired {
		t.Errorf(
			"expected %v, got %v",
			apperror.ErrUsernameRequired,
			err,
		)
	}
}
