package usecase

import (
	"context"
	"testing"

	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAIService struct {
	mock.Mock
}

func (m *MockAIService) GenerateInsight(ctx context.Context, data string) (string, error) {
	args := m.Called(ctx, data)
	return args.String(0), args.Error(1)
}

func TestInsightUseCase_GetFinancialInsight(t *testing.T) {
	ai := new(MockAIService)
	txRepo := new(MockTransactionRepository)
	useCase := NewInsightUseCase(ai, txRepo)

	userID := "user-123"
	txRepo.On("ListByUserID", mock.Anything, userID).Return([]*entity.Transaction{}, nil)
	ai.On("GenerateInsight", mock.Anything, mock.Anything).Return("Economize mais!", nil)

	res, err := useCase.GetFinancialInsight(context.Background(), userID)

	assert.Nil(t, err)
	assert.Equal(t, "Economize mais!", res)
}
