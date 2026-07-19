package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	GitHubAPI string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env")
	}

	return Config{
		Port:      os.Getenv("PORT"),
		GitHubAPI: os.Getenv("GITHUB_API_URL"),
	}
}
