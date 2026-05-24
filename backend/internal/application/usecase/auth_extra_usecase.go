package usecase

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
)

func (u *AuthUseCase) RefreshToken(ctx context.Context, tokenStr string) (*dto.AuthResponse, error) {
	// Em produção, aqui validaríamos a assinatura do refresh token
	// Por agora, assumimos que se o token existe no banco/cache, geramos novos
	
	// Mock logic: Para o MVP, aceitamos qualquer token não vazio
	if tokenStr == "" {
		return nil, errors.New("empty refresh token")
	}

	// Simulamos a busca do usuário pelo sub do token (claims)
	// Como não temos a secret injetada nesta função específica via parâmetro global (está no service),
	// chamamos o jwtService.
	
	// Implementação real exigiria expor a secret ou ter um Validate no jwtService.
	return &dto.AuthResponse{
		AccessToken:  "new_access_token_simulated",
		RefreshToken: "new_refresh_token_simulated",
	}, nil
}
