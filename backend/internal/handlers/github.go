package handlers

import (
	"github-insights/internal/apperror"
	"github-insights/internal/models"

	"errors"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetUser(username string) (models.GitHubUser, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(
	service UserService,
) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetGitHubUser(c *fiber.Ctx) error {
	username := c.Query("user")

	user, err := h.service.GetUser(username)

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
