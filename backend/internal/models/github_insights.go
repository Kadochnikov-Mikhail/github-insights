package models

import "time"

type RepoLangs map[string]int

type GitHubInsights struct {
	Username     string    `json:"username"`
	Repositories int       `json:"repositories"`
	TotalStars   int       `json:"total_stars"`
	Languages    RepoLangs `json:"languages"`
	CreatedAt    time.Time `json:"created_at"`
}
