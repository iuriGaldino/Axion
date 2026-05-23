package usecase

import (
	"context"
	"errors"

)

func (u *AuthUseCase) RefreshToken(ctx context.Context, token string) (string, string, error) {
	// In a real scenario, we would validate the refresh token and extract the user
	// For this implementation, we return a generic pair of tokens
	return "access_token_refreshed", "refresh_token_refreshed", nil
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
