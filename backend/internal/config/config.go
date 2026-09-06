package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port        string
    GitHubAPI   string
    DatabaseURL string
}

func Load() Config {
    _ = godotenv.Load()

    return Config{
        Port:        os.Getenv("PORT"),
        GitHubAPI:   os.Getenv("GITHUB_API_URL"),
        DatabaseURL: os.Getenv("DATABASE_URL"),
    }
}