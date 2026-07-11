package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github-insights/internal/models"
)

func GetUser(username string) (models.GitHubUser, error) {
	var user models.GitHubUser
	response, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		return models.GitHubUser{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.GitHubUser{}, fmt.Errorf("%d GitHub API response status", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(&user)

	if err != nil {
		return models.GitHubUser{}, err
	}
	return user, nil
}
