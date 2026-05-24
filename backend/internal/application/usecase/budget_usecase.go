package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/iuriGaldino/Axion/backend/internal/domain/repository"
)

type BudgetUseCase struct {
	budgetRepo      repository.BudgetRepository
	transactionRepo repository.TransactionRepository
}

func NewBudgetUseCase(br repository.BudgetRepository, tr repository.TransactionRepository) *BudgetUseCase {
	return &BudgetUseCase{budgetRepo: br, transactionRepo: tr}
}

func (u *BudgetUseCase) CreateBudget(ctx context.Context, userID string, req dto.CreateBudgetRequest) error {
	uID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	cID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return err
	}
	budget := entity.NewBudget(uID, cID, req.Amount, req.Period, req.StartDate, req.EndDate)
	return u.budgetRepo.Create(ctx, budget)
}

func (u *BudgetUseCase) MonitorBudgets(ctx context.Context, userID string) (map[string]float64, error) {
	budgets, err := u.budgetRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	transactions, err := u.transactionRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	usage := make(map[string]float64)
	for _, b := range budgets {
		var totalSpent float64
		for _, t := range transactions {
			if t.CategoryID != nil && *t.CategoryID == b.CategoryID && t.Type == entity.Expense {
				if t.Date.After(b.StartDate) && (b.EndDate.IsZero() || t.Date.Before(b.EndDate)) {
					totalSpent += t.Amount
				}
			}
		}
		if b.Amount > 0 {
			usage[b.CategoryID.String()] = (totalSpent / b.Amount) * 100
		} else {
			usage[b.CategoryID.String()] = 0
		}
	}
	return usage, nil
}
