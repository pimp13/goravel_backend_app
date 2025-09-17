package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type JwtRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *JwtRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *JwtRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *JwtRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *JwtRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *JwtRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *JwtRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
