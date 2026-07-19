package models

type RepoLangs map[string]int

type GitHubInsights struct {
	Username     string    `json:"username"`
	Repositories int       `json:"repositories"`
	TotalStars   int       `json:"total_stars"`
	Languages    RepoLangs `json:"languages"`
}
