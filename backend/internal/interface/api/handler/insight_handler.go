package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iuriGaldino/Axion/backend/internal/application/usecase"
)

type InsightHandler struct {
	useCase *usecase.InsightUseCase
}

func NewInsightHandler(u *usecase.InsightUseCase) *InsightHandler {
	return &InsightHandler{useCase: u}
}

func (h *InsightHandler) GetInsight(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	insight, err := h.useCase.GetFinancialInsight(c.Context(), userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"insight": insight})
}
