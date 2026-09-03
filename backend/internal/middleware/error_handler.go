package middleware

import (
	"errors"

	"github-insights/internal/apperror"
	"github-insights/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	switch {
	case errors.Is(err, apperror.ErrUsernameRequired):
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, apperror.ErrUserNotFound):
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: err.Error(),
		})

	default:
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}
}
