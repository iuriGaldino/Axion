package usecase

import (
	"context"
	"testing"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/google/uuid"
)

func TestTransactionUseCase_List(t *testing.T) {
	repo := new(MockTransactionRepository)
	accRepo := new(MockAccountRepository)
	useCase := NewTransactionUseCase(repo, accRepo)

	userID := uuid.New().String()
	repo.On("ListByUserID", mock.Anything, userID).Return([]*entity.Transaction{{Amount: 10.0}}, nil)

	res, err := useCase.List(context.Background(), userID)

	assert.Nil(t, err)
	assert.Len(t, res, 1)
}
