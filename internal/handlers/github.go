package handlers

import (
	"github-insights/internal/apperror"
	"github-insights/internal/models"
	"github-insights/internal/services"

	"errors"
	"github.com/gofiber/fiber/v2"
)

func GetGitHubUser(c *fiber.Ctx) error {
	username := c.Query("user")
	user, err := services.GetUser(username)
	if err != nil {
		if errors.Is(err, apperror.ErrUsernameRequired) {
			return c.Status(fiber.StatusBadRequest).JSON(
				models.ErrorResponse{
					Error: err.Error(),
				},
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			models.ErrorResponse{
				Error: err.Error(),
			},
		)
	}

	return c.JSON(user)
}
