package models

import "github.com/goravel/framework/database/orm"

type Tag struct {
	orm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`

	Posts []*Post `json:"posts" gorm:"many2many:post_tag;foreignKey:ID;joinForeignKey:TagID;joinReferences:PostID"`
}
