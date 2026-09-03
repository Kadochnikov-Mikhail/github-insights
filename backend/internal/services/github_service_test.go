package services

import (
	"errors"
	"testing"

	"github-insights/internal/apperror"
	"github-insights/internal/models"
)

type MockGitHubClient struct {
	User  models.GitHubUser
	Repos []models.GitHubRepo
	Err   error
}

func (m MockGitHubClient) GetUser(username string) (models.GitHubUser, error) {
	return m.User, m.Err
}

func (m MockGitHubClient) GetUserRepos(username string) ([]models.GitHubRepo, error) {
	return m.Repos, m.Err
}

func TestGetUserSuccess(t *testing.T) {

	client := MockGitHubClient{
		User: models.GitHubUser{
			Login: "torvalds",
		},
	}

	service := NewUserService(client)

	user, err := service.GetUser("torvalds")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Login != "torvalds" {
		t.Errorf("expected torvalds, got %s", user.Login)
	}
}

func TestGetUserClientError(t *testing.T) {

	client := MockGitHubClient{
		Err: errors.New("github api error"),
	}

	service := NewUserService(client)

	_, err := service.GetUser("torvalds")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetUserEmptyUsername(t *testing.T) {

	client := MockGitHubClient{}

	service := NewUserService(client)

	_, err := service.GetUser("")

	if err != apperror.ErrUsernameRequired {
		t.Errorf(
			"expected %v, got %v",
			apperror.ErrUsernameRequired,
			err,
		)
	}
}
