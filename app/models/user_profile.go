package models

import "github.com/goravel/framework/database/orm"

type UserProfile struct {
	orm.Model
	Location string `json:"location,omitempty"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`

	UserId uint  `json:"user_id"`
	User   *User `json:"user" gorm:"foreignKey:UserId"`
}
