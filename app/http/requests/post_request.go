package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type PostRequest struct {
	Title      string   `json:"title" form:"title"`
	Summary    string   `json:"summary" form:"summary"`
	Content    string   `json:"content" form:"content"`
	IsActive   bool     `json:"is_active" form:"is_active"`
	Slug       string   `json:"slug" form:"slug"`
	CategoryId uint     `json:"category_id" form:"category_id"`
	Tags       []string `json:"tags" form:"tags"`
}

func (r *PostRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *PostRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PostRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"title":       "required|max_len:195|string",
		"slug":        "required|max_len:175|string|unique:posts,slug",
		"summary":     "string",
		"content":     "required",
		"image":       "required|image",
		"category_id": "required",
		"is_active":   "required|bool",
		"tags.*":      "unique:tags,name",
	}
}

func (r *PostRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PostRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PostRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
