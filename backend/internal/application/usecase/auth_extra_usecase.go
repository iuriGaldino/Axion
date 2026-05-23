package usecase

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

func (u *AuthUseCase) RefreshToken(ctx context.Context, tokenStr string) (string, string, error) {
	// Validação real do token seria feita aqui com a secret
	return "new_access_token", "new_refresh_token", nil
}

func (u *AuthUseCase) ChangePassword(ctx context.Context, userID, oldPass, newPass string) error {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if !u.hashService.Compare(user.PasswordHash, oldPass) {
		return errors.New("invalid old password")
	}
	hashed, err := u.hashService.Hash(newPass)
	if err != nil {
		return err
	}
	user.PasswordHash = hashed
	return u.userRepo.Update(ctx, user)
}
