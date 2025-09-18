package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type TagRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *TagRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *TagRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *TagRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name": "required|string|unique:tags,name",
	}
}

func (r *TagRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *TagRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *TagRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
