package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/application/usecase"
)

type CategoryHandler struct {
	useCase *usecase.CategoryUseCase
}

func NewCategoryHandler(u *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{useCase: u}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req dto.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	res, err := h.useCase.Create(c.Context(), userID, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(res)
}

func (h *CategoryHandler) List(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	res, err := h.useCase.List(c.Context(), userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(res)
}
