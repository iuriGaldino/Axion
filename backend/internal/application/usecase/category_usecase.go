package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/iuriGaldino/Axion/backend/internal/application/dto"
	"github.com/iuriGaldino/Axion/backend/internal/domain/entity"
	"github.com/iuriGaldino/Axion/backend/internal/domain/repository"
)

type CategoryUseCase struct {
	repo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{repo: repo}
}

func (u *CategoryUseCase) Create(ctx context.Context, userID string, req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	uID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	category := entity.NewCategory(uID, req.Name, req.Icon, req.Color, req.IsIncome)
	if err := u.repo.Create(ctx, category); err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:       category.ID.String(),
		Name:     category.Name,
		Icon:     category.Icon,
		Color:    category.Color,
		IsIncome: category.IsIncome,
	}, nil
}

func (u *CategoryUseCase) List(ctx context.Context, userID string) ([]*dto.CategoryResponse, error) {
	categories, err := u.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]*dto.CategoryResponse, len(categories))
	for i, c := range categories {
		res[i] = &dto.CategoryResponse{
			ID:       c.ID.String(),
			Name:     c.Name,
			Icon:     c.Icon,
			Color:    c.Color,
			IsIncome: c.IsIncome,
		}
	}
	return res, nil
}
