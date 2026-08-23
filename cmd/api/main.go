package main

import (
	"log"

	"github-insights/internal/client"
	"github-insights/internal/config"
	"github-insights/internal/database"
	"github-insights/internal/handlers"
	"github-insights/internal/middleware"
	"github-insights/internal/routes"
	"github-insights/internal/services"
	"github-insights/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	insightsRepo := repository.NewInsightsRepository(db)

	defer db.Close()

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
		insightsRepo,
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

	log.Fatal(app.Listen(":" + cfg.Port))
}
