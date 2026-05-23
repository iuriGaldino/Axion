package usecase

import (
	"context"
	"errors"

)

func (u *AuthUseCase) RefreshToken(ctx context.Context, token string) (string, string, error) {
	// Implementação simplificada para o scaffold
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
	hashed, _ := u.hashService.Hash(newPass)
	user.PasswordHash = hashed
	return u.userRepo.Update(ctx, user)
}
