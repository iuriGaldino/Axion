package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(ctx context.Context, t *entity.Transaction) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetByID(ctx context.Context, id string) (*entity.Transaction, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Transaction, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) Update(ctx context.Context, t *entity.Transaction) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *MockTransactionRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestTransactionUseCase_Create(t *testing.T) {
	repo := new(MockTransactionRepository)
	accRepo := new(MockAccountRepository)
	useCase := NewTransactionUseCase(repo, accRepo)

	userID := uuid.New().String()
	accID := uuid.New().String()
	req := dto.CreateTransactionRequest{
		AccountID:   accID,
		Amount:      100.50,
		Description: "Lunch",
		Date:        time.Now(),
		Type:        "expense",
	}

	accRepo.On("GetByID", mock.Anything, accID).Return(&entity.Account{ID: uuid.MustParse(accID)}, nil)
	repo.On("Create", mock.Anything, mock.Anything).Return(nil)
	accRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

	res, err := useCase.Create(context.Background(), userID, req)

	assert.Nil(t, err)
	assert.Equal(t, 100.50, res.Amount)
}
