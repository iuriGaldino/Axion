package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/application/usecase"
	"github.com/iuriGaldino/Axion/backend/internal/infrastructure/security"
)

type AccountHandler struct {
	useCase *usecase.AccountUseCase
}

func NewAccountHandler(u *usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{useCase: u}
}

func (h *AccountHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req dto.CreateAccountRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := security.Validate(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := h.useCase.Create(c.Context(), userID, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(res)
}

func (h *AccountHandler) List(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	res, err := h.useCase.List(c.Context(), userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
