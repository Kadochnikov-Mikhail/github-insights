package handlers

import (
	"github-insights/internal/models"

	"github.com/gofiber/fiber/v2"
)

type InsightsService interface {
	GetInsights(username string) (models.GitHubInsights, error)
}

type InsightsHandler struct {
	service InsightsService
}

func NewInsightsHandler(service InsightsService) *InsightsHandler {
	return &InsightsHandler{
		service: service,
	}
}

func (h *InsightsHandler) GetInsights(c *fiber.Ctx) error {
	username := c.Query("user")

	insights, err := h.service.GetInsights(username)
	if err != nil {
		return err
	}

	return c.JSON(insights)
}
