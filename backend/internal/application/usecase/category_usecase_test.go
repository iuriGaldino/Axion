package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(ctx context.Context, c *entity.Category) error {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Category, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(ctx context.Context, c *entity.Category) error {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCategoryUseCase_Create(t *testing.T) {
	repo := new(MockCategoryRepository)
	useCase := NewCategoryUseCase(repo)

	userID := uuid.New().String()
	req := dto.CreateCategoryRequest{
		Name:     "Food",
		Icon:     "fastfood",
		Color:    "#FF0000",
		IsIncome: false,
	}

	repo.On("Create", mock.Anything, mock.Anything).Return(nil)

	res, err := useCase.Create(context.Background(), userID, req)

	assert.Nil(t, err)
	assert.Equal(t, "Food", res.Name)
	repo.AssertExpectations(t)
}
