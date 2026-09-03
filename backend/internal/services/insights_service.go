package services

import (
	"github-insights/internal/apperror"
	"github-insights/internal/client"
	"github-insights/internal/models"
	"github-insights/internal/repository"
)

type InsightsService struct {
	client     client.GitHubClient
	repository repository.InsightsRepository
}

func NewInsightsService(
	client client.GitHubClient,
	repository repository.InsightsRepository,
) *InsightsService {
	return &InsightsService{
		client:     client,
		repository: repository,
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

	insights := models.GitHubInsights{
		Username:     username,
		Repositories: len(repos),
		TotalStars:   totalStars,
		Languages:    languages,
	}
	err = s.repository.Save(insights)

	if err != nil {
		return models.GitHubInsights{}, err
	}

	return insights, nil
}

func (s *InsightsService) GetInsightsHistory(
	username string,
) ([]models.GitHubInsights, error) {

	if username == "" {
		return nil, apperror.ErrUsernameRequired
	}

	return s.repository.GetByUsername(username)
}
