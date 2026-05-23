package usecase

import (
	"context"
	"testing"

	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

type MockHashService struct {
	mock.Mock
}

func (m *MockHashService) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockHashService) Compare(hash, password string) bool {
	args := m.Called(hash, password)
	return args.Bool(0)
}

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateTokens(user *entity.User) (string, string, error) {
	args := m.Called(user)
	return args.String(0), args.String(1), args.Error(2)
}

func TestAuthUseCase_Register(t *testing.T) {
	repo := new(MockUserRepository)
	hash := new(MockHashService)
	jwt := new(MockJWTService)
	useCase := NewAuthUseCase(repo, hash, jwt)

	req := dto.RegisterRequest{
		Name:     "Test",
		Email:    "test@example.com",
		Password: "password123",
	}

	repo.On("GetByEmail", mock.Anything, req.Email).Return(nil, nil)
	hash.On("Hash", req.Password).Return("hashed", nil)
	repo.On("Create", mock.Anything, mock.Anything).Return(nil)
	jwt.On("GenerateTokens", mock.Anything).Return("access", "refresh", nil)

	res, err := useCase.Register(context.Background(), req)

	assert.Nil(t, err)
	assert.Equal(t, "access", res.AccessToken)
	repo.AssertExpectations(t)
}
