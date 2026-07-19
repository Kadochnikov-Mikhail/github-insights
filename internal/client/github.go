package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github-insights/internal/apperror"
	"github-insights/internal/models"
)

type GitHubAPIClient struct {
}

func NewGitHubAPIClient() *GitHubAPIClient {
	return &GitHubAPIClient{}
}

func (c *GitHubAPIClient) GetUser(username string) (models.GitHubUser, error) {
	var user models.GitHubUser
	response, err := httpClient.Get("https://api.github.com/users/" + username)
	if err != nil {
		return models.GitHubUser{}, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return models.GitHubUser{}, apperror.ErrUserNotFound
	}

	if response.StatusCode != http.StatusOK {
		return models.GitHubUser{}, fmt.Errorf(
			"github API returned status: %d",
			response.StatusCode,
		)
	}

	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return models.GitHubUser{}, err
	}
	return user, nil
}

func (c *GitHubAPIClient) GetUserRepos(username string) ([]models.GitHubRepo, error) {

	var repos []models.GitHubRepo

	url := fmt.Sprintf(
		"https://api.github.com/users/%s/repos",
		username,
	)

	response, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, apperror.ErrUserNotFound
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"github API returned status: %d",
			response.StatusCode,
		)
	}

	if err := json.NewDecoder(response.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}
