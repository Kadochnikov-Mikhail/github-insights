package routes

import (
	"github-insights/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/health", handlers.Health)
	app.Get("/github", handlers.GetGitHubUser)
	app.Get("/github/insights", handlers.GetInsights)
}
