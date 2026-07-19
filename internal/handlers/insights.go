package handlers

import (
	"errors"

	"github-insights/internal/apperror"
	"github-insights/internal/models"
	"github-insights/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetInsights(c *fiber.Ctx) error {
	username := c.Query("user")

	insights, err := services.GetInsights(username)
	if err != nil {

		switch {
		case errors.Is(err, apperror.ErrUsernameRequired):
			return c.Status(400).JSON(models.ErrorResponse{
				Error: err.Error(),
			})

		case errors.Is(err, apperror.ErrUserNotFound):
			return c.Status(404).JSON(models.ErrorResponse{
				Error: err.Error(),
			})

		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Error: err.Error(),
			})
		}
	}

	return c.JSON(insights)
}
