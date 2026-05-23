package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/iuriGaldino/Axion/backend/internal/domain/repository"
)

type AccountUseCase struct {
	repo repository.AccountRepository
}

func NewAccountUseCase(repo repository.AccountRepository) *AccountUseCase {
	return &AccountUseCase{repo: repo}
}

func (u *AccountUseCase) Create(ctx context.Context, userID string, req dto.CreateAccountRequest) (*dto.AccountResponse, error) {
	uID, _ := uuid.Parse(userID)
	account := entity.NewAccount(uID, req.Name, req.Type, req.Currency)
	if err := u.repo.Create(ctx, account); err != nil {
		return nil, err
	}
	return &dto.AccountResponse{
		ID: account.ID.String(),
		Name: account.Name,
		Type: account.Type,
		Balance: account.Balance,
		Currency: account.Currency,
	}, nil
}

func (u *AccountUseCase) List(ctx context.Context, userID string) ([]*dto.AccountResponse, error) {
	accounts, err := u.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := make([]*dto.AccountResponse, len(accounts))
	for i, a := range accounts {
		res[i] = &dto.AccountResponse{
			ID: a.ID.String(),
			Name: a.Name,
			Type: a.Type,
			Balance: a.Balance,
			Currency: a.Currency,
		}
	}
	return res, nil
}
