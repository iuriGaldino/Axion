package usecase

import (
	"context"
	"testing"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthUseCase_Login(t *testing.T) {
	repo := new(MockUserRepository)
	hash := new(MockHashService)
	jwt := new(MockJWTService)
	useCase := NewAuthUseCase(repo, hash, jwt)

	user := entity.NewUser("Test", "test@example.com", "hashed")
	req := dto.LoginRequest{Email: "test@example.com", Password: "password123"}

	repo.On("GetByEmail", mock.Anything, req.Email).Return(user, nil)
	hash.On("Compare", "hashed", req.Password).Return(true)
	jwt.On("GenerateTokens", mock.Anything).Return("access", "refresh", nil)

	res, err := useCase.Login(context.Background(), req)

	assert.Nil(t, err)
	assert.Equal(t, "access", res.AccessToken)
}
