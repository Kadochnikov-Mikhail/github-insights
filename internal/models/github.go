package models

type GitHubUser struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

type GitHubRepo struct {
	Name     string `json:"name"`
	Stars    int    `json:"stargazers_count"`
	Language string `json:"language"`
}
