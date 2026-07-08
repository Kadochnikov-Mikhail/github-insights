package handlers

import (
	"github-insights/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(models.HealthResponse{
		Status: "ok",
	})
}
