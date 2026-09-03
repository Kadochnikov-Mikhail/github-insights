package models

type GitHubRepo struct {
	Name     string `json:"name"`
	Stars    int    `json:"stargazers_count"`
	Language string `json:"language"`
}
