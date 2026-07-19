package main

import (
	"github-insights/internal/client"
	"github-insights/internal/config"
	"github-insights/internal/handlers"
	"github-insights/internal/middleware"
	"github-insights/internal/routes"
	"github-insights/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.Load()

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(middleware.Logger)

	githubClient := client.NewGitHubAPIClient()

	userService := services.NewUserService(
		githubClient,
	)

	insightsService := services.NewInsightsService(
		githubClient,
	)

	userHandler := handlers.NewUserHandler(
		userService,
	)

	insightsHandler := handlers.NewInsightsHandler(
		insightsService,
	)

	routes.Setup(
		app,
		userHandler,
		insightsHandler,
	)

	if err := app.Listen(":" + cfg.Port); err != nil {
		panic(err)
	}
}
