package usecase

import (
	"context"
	"errors"

	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/iuriGaldino/Axion/backend/internal/domain/repository"
)

type AuthUseCase struct {
	userRepo    repository.UserRepository
	hashService interface {
		Hash(password string) (string, error)
		Compare(hash, password string) bool
	}
	jwtService interface {
		GenerateTokens(user *entity.User) (string, string, error)
	}
}

func NewAuthUseCase(
	userRepo repository.UserRepository,
	hashService interface {
		Hash(password string) (string, error)
		Compare(hash, password string) bool
	},
	jwtService interface {
		GenerateTokens(user *entity.User) (string, string, error)
	},
) *AuthUseCase {
	return &AuthUseCase{
		userRepo:    userRepo,
		hashService: hashService,
		jwtService:  jwtService,
	}
}

func (u *AuthUseCase) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existing, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existing != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := u.hashService.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(req.Name, req.Email, hashedPassword)
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	access, refresh, err := u.jwtService.GenerateTokens(user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (u *AuthUseCase) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !u.hashService.Compare(user.PasswordHash, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	access, refresh, err := u.jwtService.GenerateTokens(user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
