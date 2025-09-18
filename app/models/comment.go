package models

import "github.com/goravel/framework/database/orm"

type Comment struct {
	orm.Model
	Content  string `json:"content"`
	IsActive bool   `json:"is_active"`
	UserId   uint   `json:"user_id"`
	PostId   uint   `json:"post_id"`

	User *User `json:"user" gorm:"foreignKey:UserId"`
	Post *Post `json:"post" gorm:"foreignKey:PostId"`
}
