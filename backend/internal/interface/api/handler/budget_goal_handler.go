package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/application/usecase"
)

type BudgetGoalHandler struct {
	useCase *usecase.BudgetGoalUseCase
}

func NewBudgetGoalHandler(u *usecase.BudgetGoalUseCase) *BudgetGoalHandler {
	return &BudgetGoalHandler{useCase: u}
}

func (h *BudgetGoalHandler) CreateGoal(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req dto.CreateGoalRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	res, err := h.useCase.CreateGoal(c.Context(), userID, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(res)
}

func (h *BudgetGoalHandler) ListGoals(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	res, err := h.useCase.ListGoals(c.Context(), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(res)
}
