package service

import (
	"context"
)

type AIService interface {
	GenerateInsight(ctx context.Context, data string) (string, error)
}

type OCRService interface {
	ExtractText(imagePath string) (string, error)
}
