package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CategoryRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *CategoryRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CategoryRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CategoryRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
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
