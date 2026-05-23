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

	if len(transactions) == 0 {
		return "Ainda não tenho dados suficientes para gerar insights. Comece adicionando algumas transações!", nil
	}

	data := "Histórico de Transações do Usuário: "
	for _, t := range transactions {
		data += fmt.Sprintf("Valor: %.2f, Descrição: %s, Tipo: %s; ", t.Amount, t.Description, t.Type)
	}

	return u.aiService.GenerateInsight(ctx, data)
}
