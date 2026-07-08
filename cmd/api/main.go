package main

import (
	"github-insights/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
