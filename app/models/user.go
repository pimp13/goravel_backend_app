package models

import (
	"goravel_by_gin/app/enums"

	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	Password      string             `json:"password"`
	IsActive      bool               `json:"is_active"`
	IsSupperAdmin bool               `json:"is_supperadmin" gorm:"column:is_supperadmin"`
	Role          enums.UserRoleEnum `json:"role"`

	Profile   *UserProfile `json:"profile" gorm:"foreignKey:UserId"`
	Posts     []*Post      `json:"posts" gorm:"foreignKey:UserId"`
	Comments  []*Comment   `json:"comments" gorm:"foreignKey:UserId"`
	Followers []*User      `gorm:"many2many:user_followers;foreignKey:ID;joinForeignKey:FollowedID;joinReferences:FollowerID"`
	Following []*User      `gorm:"many2many:user_followers;foreignKey:ID;joinForeignKey:FollowerID;joinReferences:FollowedID"`
}
