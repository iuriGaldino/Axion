package repository

import (
	"context"

	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
)

type BudgetRepository interface {
	Create(ctx context.Context, budget *entity.Budget) error
	ListByUserID(ctx context.Context, userID string) ([]*entity.Budget, error)
	Delete(ctx context.Context, id string) error
}

type GoalRepository interface {
	Create(ctx context.Context, goal *entity.Goal) error
	GetByID(ctx context.Context, id string) (*entity.Goal, error)
	ListByUserID(ctx context.Context, userID string) ([]*entity.Goal, error)
	Update(ctx context.Context, goal *entity.Goal) error
	Delete(ctx context.Context, id string) error
}
