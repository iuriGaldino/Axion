package usecase

import (
	"context"
	"testing"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/google/uuid"
)

type MockAccountRepository struct {
	mock.Mock
}
func (m *MockAccountRepository) Create(ctx context.Context, a *entity.Account) error { return m.Called(ctx, a).Error(0) }
func (m *MockAccountRepository) GetByID(ctx context.Context, id string) (*entity.Account, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*entity.Account), args.Error(1)
}
func (m *MockAccountRepository) ListByUserID(ctx context.Context, userID string) ([]*entity.Account, error) {
	return m.Called(ctx, userID).Get(0).([]*entity.Account), m.Called(ctx, userID).Error(1)
}
func (m *MockAccountRepository) Update(ctx context.Context, a *entity.Account) error { return m.Called(ctx, a).Error(0) }
func (m *MockAccountRepository) Delete(ctx context.Context, id string) error { return m.Called(ctx, id).Error(0) }
