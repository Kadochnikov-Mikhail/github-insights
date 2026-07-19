package services

import (
	"github-insights/internal/apperror"
	"github-insights/internal/client"
	"github-insights/internal/models"
)

type UserService struct {
	client client.GitHubClient
}

func NewUserService(
	client client.GitHubClient,
) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) GetUser(username string) (models.GitHubUser, error) {

	if username == "" {
		return models.GitHubUser{}, apperror.ErrUsernameRequired
	}

	return s.client.GetUser(username)
}
