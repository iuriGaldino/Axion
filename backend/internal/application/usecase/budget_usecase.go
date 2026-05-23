package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

func (u *BudgetGoalUseCase) CreateBudget(ctx context.Context, userID string, req dto.CreateBudgetRequest) error {
	uID, _ := uuid.Parse(userID)
	cID, _ := uuid.Parse(req.CategoryID)
	budget := entity.NewBudget(uID, cID, req.Amount, req.Period, req.StartDate, req.EndDate)
	return u.budgetRepo.Create(ctx, budget)
}

func (u *BudgetGoalUseCase) MonitorBudgets(ctx context.Context, userID string) (map[string]float64, error) {
	// Lógica para monitorar gastos vs orçamentos
	return make(map[string]float64), nil
}
