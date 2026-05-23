package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/application/usecase"
	"github.com/iuriGaldino/Axion/backend/internal/infrastructure/security"
)

type BudgetHandler struct {
	useCase *usecase.BudgetUseCase
}

func NewBudgetHandler(u *usecase.BudgetUseCase) *BudgetHandler {
	return &BudgetHandler{useCase: u}
}

func (h *BudgetHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req dto.CreateBudgetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := security.Validate(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.useCase.CreateBudget(c.Context(), userID, req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(201)
}

func (h *BudgetHandler) Monitor(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	res, err := h.useCase.MonitorBudgets(c.Context(), userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
