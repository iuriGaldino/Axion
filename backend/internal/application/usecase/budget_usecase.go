package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

func (u *BudgetGoalUseCase) CreateBudget(ctx context.Context, userID string, req dto.CreateBudgetRequest) error {
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

func (u *BudgetGoalUseCase) MonitorBudgets(ctx context.Context, userID string) (map[string]float64, error) {
	budgets, err := u.budgetRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	usage := make(map[string]float64)
	for _, b := range budgets {
		// Em produção, aqui haveria uma query para somar transações do período
		usage[b.CategoryID.String()] = 0.0
	}
	return usage, nil
}
