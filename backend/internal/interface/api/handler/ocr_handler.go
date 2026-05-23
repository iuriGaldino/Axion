package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iuriGaldino/Axion/backend/internal/domain/service"
)

type OCRHandler struct {
	ocrService service.OCRService
}

func NewOCRHandler(s service.OCRService) *OCRHandler {
	return &OCRHandler{ocrService: s}
}

func (h *OCRHandler) ProcessReceipt(c *fiber.Ctx) error {
	// Simulação de upload e processamento
	text, err := h.ocrService.ExtractText("receipt.jpg")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"text": text})
}
