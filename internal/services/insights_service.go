package services

import (
	"github-insights/internal/apperror"
	"github-insights/internal/client"
	"github-insights/internal/models"
)

type InsightsService struct {
	client client.GitHubClient
}

func NewInsightsService(client client.GitHubClient) *InsightsService {
	return &InsightsService{
		client: client,
	}
}

func (s *InsightsService) GetInsights(username string) (models.GitHubInsights, error) {

	if username == "" {
		return models.GitHubInsights{}, apperror.ErrUsernameRequired
	}

	repos, err := s.client.GetUserRepos(username)
	if err != nil {
		return models.GitHubInsights{}, err
	}

	totalStars := 0
	languages := make(models.RepoLangs)

	for _, repo := range repos {
		totalStars += repo.Stars

		if repo.Language != "" {
			languages[repo.Language]++
		}
	}

	return models.GitHubInsights{
		Username:     username,
		Repositories: len(repos),
		TotalStars:   totalStars,
		Languages:    languages,
	}, nil
}
