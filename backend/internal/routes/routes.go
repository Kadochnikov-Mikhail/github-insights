package routes

import (
	"github-insights/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(
	app *fiber.App,
	userHandler *handlers.UserHandler,
	insightsHandler *handlers.InsightsHandler,
) {
	app.Get("/health", handlers.Health)
	app.Get("/github", userHandler.GetGitHubUser)
	app.Get("/github/insights", insightsHandler.GetInsights)
	app.Get("/github/insights/history", insightsHandler.GetInsightsHistory)
}
