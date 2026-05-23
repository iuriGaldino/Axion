package usecase

import (
	"context"
	"fmt"
	"github.com/iuriGaldino/Axion/backend/internal/domain/repository"
	"github.com/iuriGaldino/Axion/backend/internal/domain/service"
)

type InsightUseCase struct {
	aiService service.AIService
	txRepo    repository.TransactionRepository
}

func NewInsightUseCase(ai service.AIService, tx repository.TransactionRepository) *InsightUseCase {
	return &InsightUseCase{aiService: ai, txRepo: tx}
}

func (u *InsightUseCase) GetFinancialInsight(ctx context.Context, userID string) (string, error) {
	transactions, err := u.txRepo.ListByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	data := "Transações recentes: "
	for _, t := range transactions {
		data += fmt.Sprintf("%.2f em %s; ", t.Amount, t.Description)
	}

	return u.aiService.GenerateInsight(ctx, data)
}
