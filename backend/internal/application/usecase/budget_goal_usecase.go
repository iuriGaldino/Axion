package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/iuriGaldino/Axion/backend/internal/domain/repository"
)

type BudgetGoalUseCase struct {
	budgetRepo repository.BudgetRepository
	goalRepo   repository.GoalRepository
}

func NewBudgetGoalUseCase(br repository.BudgetRepository, gr repository.GoalRepository) *BudgetGoalUseCase {
	return &BudgetGoalUseCase{budgetRepo: br, goalRepo: gr}
}

func (u *BudgetGoalUseCase) CreateGoal(ctx context.Context, userID string, req dto.CreateGoalRequest) (*dto.GoalResponse, error) {
	uID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	goal := entity.NewGoal(uID, req.Name, req.TargetAmount, req.Deadline)
	if err := u.goalRepo.Create(ctx, goal); err != nil {
		return nil, err
	}

	return &dto.GoalResponse{
		ID:            goal.ID.String(),
		Name:          goal.Name,
		TargetAmount:  goal.TargetAmount,
		CurrentAmount: goal.CurrentAmount,
		Progress:      goal.CalculateProgress(),
		Deadline:      goal.Deadline,
	}, nil
}

func (u *BudgetGoalUseCase) ListGoals(ctx context.Context, userID string) ([]*dto.GoalResponse, error) {
	goals, err := u.goalRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]*dto.GoalResponse, len(goals))
	for i, g := range goals {
		res[i] = &dto.GoalResponse{
			ID:            g.ID.String(),
			Name:          g.Name,
			TargetAmount:  g.TargetAmount,
			CurrentAmount: g.CurrentAmount,
			Progress:      g.CalculateProgress(),
			Deadline:      g.Deadline,
		}
	}
	return res, nil
}
