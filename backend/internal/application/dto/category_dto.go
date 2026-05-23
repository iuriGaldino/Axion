package dto

type CreateCategoryRequest struct {
	Name     string `json:"name" validate:"required"`
	Icon     string `json:"icon"`
	Color    string `json:"color"`
	IsIncome bool   `json:"is_income"`
}

type CategoryResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Color    string `json:"color"`
	IsIncome bool   `json:"is_income"`
}
