package services

import (
	"errors"

	"github-insights/internal/client"
	"github-insights/internal/models"
)

var ErrUsernameRequired = errors.New("username is required")

func GetUser(username string) (models.GitHubUser, error) {
	if username == "" {
		return models.GitHubUser{}, ErrUsernameRequired
	}

	return client.GetUser(username)
}
