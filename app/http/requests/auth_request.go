package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *AuthRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{
		"email": "trim",
		"name":  "trim",
	}
}

func (r *AuthRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required|string|max_len:175|min_len:4",
		"email":    "required|email|max_len:175|min_len:5|unique:users,email",
		"password": "required|min_len:8",
	}
}

func (r *AuthRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}

type AuthLoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *AuthLoginRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthLoginRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{
		"email": "trim",
		"name":  "trim",
	}
}

func (r *AuthLoginRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email|max_len:175|min_len:5",
		"password": "required|min_len:8",
	}
}

func (r *AuthLoginRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthLoginRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthLoginRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
