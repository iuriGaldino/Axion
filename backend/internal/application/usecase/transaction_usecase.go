package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/iuriGaldino/Axion/backend/internal/domain/repository"
)

type TransactionUseCase struct {
	repo        repository.TransactionRepository
	accountRepo repository.AccountRepository
}

func NewTransactionUseCase(repo repository.TransactionRepository, accountRepo repository.AccountRepository) *TransactionUseCase {
	return &TransactionUseCase{repo: repo, accountRepo: accountRepo}
}

func (u *TransactionUseCase) Create(ctx context.Context, userID string, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
	// Em produção, envolveríamos tudo em um db.BeginTx
	uID, _ := uuid.Parse(userID)
	aID, err := uuid.Parse(req.AccountID)
	if err != nil {
		return nil, errors.New("invalid account id")
	}

	account, err := u.accountRepo.GetByID(ctx, req.AccountID)
	if err != nil {
		return nil, errors.New("account not found")
	}

	var cID *uuid.UUID
	if req.CategoryID != "" {
		parsedCID, err := uuid.Parse(req.CategoryID)
		if err == nil {
			cID = &parsedCID
		}
	}

	transaction := entity.NewTransaction(uID, aID, cID, req.Amount, req.Description, req.Date, entity.TransactionType(req.Type))
	
	if transaction.Type == entity.Income {
		account.Balance += transaction.Amount
	} else if transaction.Type == entity.Expense {
		account.Balance -= transaction.Amount
	}

	if err := u.repo.Create(ctx, transaction); err != nil {
		return nil, err
	}

	if err := u.accountRepo.Update(ctx, account); err != nil {
		return nil, err
	}

	return &dto.TransactionResponse{
		ID:          transaction.ID.String(),
		AccountID:   transaction.AccountID.String(),
		CategoryID:  req.CategoryID,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Date:        transaction.Date,
		Type:        string(transaction.Type),
		IsConfirmed: transaction.IsConfirmed,
	}, nil
}

func (u *TransactionUseCase) List(ctx context.Context, userID string) ([]*dto.TransactionResponse, error) {
	transactions, err := u.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]*dto.TransactionResponse, len(transactions))
	for i, t := range transactions {
		catID := ""
		if t.CategoryID != nil {
			catID = t.CategoryID.String()
		}
		res[i] = &dto.TransactionResponse{
			ID:          t.ID.String(),
			AccountID:   t.AccountID.String(),
			CategoryID:  catID,
			Amount:      t.Amount,
			Description: t.Description,
			Date:        t.Date,
			Type:        string(t.Type),
			IsConfirmed: t.IsConfirmed,
		}
	}
	return res, nil
}
