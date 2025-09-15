package models

import "github.com/goravel/framework/database/orm"

type Category struct {
	orm.Model
	Name        string `json:"name"`
	IsActive    bool   `json:"is_active"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
