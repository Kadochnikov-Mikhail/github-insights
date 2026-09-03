package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github-insights/internal/apperror"
	"github-insights/internal/models"

	"github.com/gofiber/fiber/v2"
)

type MockUserService struct{}
type MockUserServiceError struct{}

func (m MockUserServiceError) GetUser(username string) (models.GitHubUser, error) {
	return models.GitHubUser{}, apperror.ErrUsernameRequired
}

func (m MockUserService) GetUser(username string) (models.GitHubUser, error) {
	return models.GitHubUser{
		Login: "torvalds",
	}, nil
}

func TestGetGitHubUserSuccess(t *testing.T) {

	app := fiber.New()

	service := MockUserService{}

	handler := NewUserHandler(service)

	app.Get("/github", handler.GetGitHubUser)

	req := httptest.NewRequest(
		"GET",
		"/github?user=torvalds",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var user models.GitHubUser

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if user.Login != "torvalds" {
		t.Errorf("expected torvalds, got %s", user.Login)
	}
}

func TestGetGitHubUserEmptyUsername(t *testing.T) {

	app := fiber.New()

	service := MockUserServiceError{}

	handler := NewUserHandler(service)

	app.Get("/github", handler.GetGitHubUser)

	req := httptest.NewRequest(
		"GET",
		"/github",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", resp.StatusCode)
	}
}
