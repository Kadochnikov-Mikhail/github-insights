package services

import (
	"github-insights/internal/apperror"
	"github-insights/internal/client"
	"github-insights/internal/models"
)

func GetUser(username string) (models.GitHubUser, error) {
	if username == "" {
		return models.GitHubUser{}, apperror.ErrUsernameRequired
	}

	return client.GetUser(username)
}
