package usecase

import (
	"context"
	"testing"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/google/uuid"
)

type MockBudgetRepository struct {
	mock.Mock
}
func (m *MockBudgetRepository) Create(ctx context.Context, b *entity.Budget) error { return m.Called(ctx, b).Error(0) }
func (m *MockBudgetRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Budget, error) {
	return m.Called(ctx, userID).Get(0).([]*entity.Budget), m.Called(ctx, userID).Error(1)
}
func (m *MockBudgetRepository) Delete(ctx context.Context, id string) error { return m.Called(ctx, id).Error(0) }

type MockGoalRepository struct {
	mock.Mock
}
func (m *MockGoalRepository) Create(ctx context.Context, g *entity.Goal) error { return m.Called(ctx, g).Error(0) }
func (m *MockGoalRepository) GetByID(ctx context.Context, id string) (*entity.Goal, error) { return m.Called(ctx, id).Get(0).(*entity.Goal), m.Called(ctx, id).Error(1) }
func (m *MockGoalRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Goal, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*entity.Goal), args.Error(1)
}
func (m *MockGoalRepository) Update(ctx context.Context, g *entity.Goal) error { return m.Called(ctx, g).Error(0) }
func (m *MockGoalRepository) Delete(ctx context.Context, id string) error { return m.Called(ctx, id).Error(0) }

func TestBudgetGoalUseCase_ListGoals(t *testing.T) {
	br := new(MockBudgetRepository)
	gr := new(MockGoalRepository)
	useCase := NewBudgetGoalUseCase(br, gr)

	userID := uuid.New().String()
	gr.On("ListByUserID", mock.Anything, userID).Return([]*entity.Goal{{Name: "Goal"}}, nil)

	res, err := useCase.ListGoals(context.Background(), userID)

	assert.Nil(t, err)
	assert.Len(t, res, 1)
}
