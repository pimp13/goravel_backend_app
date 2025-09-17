package models

import "github.com/goravel/framework/database/orm"

type Post struct {
	orm.Model
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	IsActive   bool   `json:"is_active"`
	Slug       string `json:"slug"`
	ImageUrl   string `json:"image_url"`
	UserId     uint   `json:"user_id"`
	CategoryId uint   `json:"category_id"`

	Category *Category `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	User     *User     `json:"user" gorm:"foreignKey:UserId"`
}
