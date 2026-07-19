package client

import "github-insights/internal/models"

type GitHubClient interface {
	GetUser(username string) (models.GitHubUser, error)
	GetUserRepos(username string) ([]models.GitHubRepo, error)
}
