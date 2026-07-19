package main

import (
	"github-insights/internal/config"
	"github-insights/internal/middleware"
	"github-insights/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.Load()

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(middleware.Logger)

	routes.Setup(app)

	if err := app.Listen(":" + cfg.Port); err != nil {
		panic(err)
	}
}
