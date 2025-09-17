package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CategoryRequest struct {
	Name        string `json:"name" form:"name"`
	Slug        string `json:"slug" form:"slug"`
	IsActive    bool   `json:"is_active" form:"is_active"`
	Description string `json:"description" form:"description"`
}

func (r *CategoryRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CategoryRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CategoryRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":        "required|min_len:5|string|max_len:195",
		"slug":        "required|min_len:3|string|max_len:175|unique:categories,slug",
		"description": "required|min_len:5|string",
		"is_active":   "required|bool",
	}
}

func (r *CategoryRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CategoryRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CategoryRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" form:"name"`
	Slug        string `json:"slug" form:"slug"`
	IsActive    bool   `json:"is_active" form:"is_active"`
	Description string `json:"description" form:"description"`
}

func (r *UpdateCategoryRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateCategoryRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateCategoryRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":        "min_len:5|string|max_len:195",
		"slug":        "min_len:3|string|max_len:175",
		"description": "min_len:5|string",
		"is_active":   "bool",
	}
}

func (r *UpdateCategoryRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateCategoryRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateCategoryRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}

type ChangeStatusCategoryDto struct {
	IsActive string `json:"is_active" form:"is_active"`
}

func (r *ChangeStatusCategoryDto) Authorize(ctx http.Context) error {
	return nil
}

func (r *ChangeStatusCategoryDto) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ChangeStatusCategoryDto) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"is_active": "required|bool",
	}
}

func (r *ChangeStatusCategoryDto) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ChangeStatusCategoryDto) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ChangeStatusCategoryDto) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
