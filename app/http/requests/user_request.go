package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *UserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required|max_len:150|min_len:5",
		"email":    "required|email",
		"password": "required|min_len:8",
	}
}

func (r *UserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
