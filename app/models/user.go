package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`

	Posts []*Post `json:"posts" gorm:"foreignKey:UserId"`
}
