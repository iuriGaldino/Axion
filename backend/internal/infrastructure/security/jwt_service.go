package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

type JWTService struct {
	secretKey     []byte
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secretKey:     []byte(secret),
		accessExpiry:  time.Hour * 1,
		refreshExpiry: time.Hour * 24 * 7,
	}
}

func (s *JWTService) GenerateTokens(user *entity.User) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"sub":   user.ID.String(),
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(s.accessExpiry).Unix(),
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(s.secretKey)
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(s.refreshExpiry).Unix(),
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(s.secretKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
