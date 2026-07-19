package handlers

import (
	"github-insights/internal/services"

	"github.com/gofiber/fiber/v2"
)

type InsightsHandler struct {
	service *services.InsightsService
}

func NewInsightsHandler(service *services.InsightsService) *InsightsHandler {
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
