package handlers

import (
	"github-insights/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetInsights(c *fiber.Ctx) error {
	username := c.Query("user")
	insights, err := services.GetInsights(username)
	if err != nil {
		return err
	}
	return c.JSON(insights)
}
