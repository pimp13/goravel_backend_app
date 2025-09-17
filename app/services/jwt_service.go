package services

import "github.com/goravel/framework/contracts/database/orm"

type JwtService interface{}

type jwtService struct {
	orm orm.Orm
}

func NewJwtService(orm orm.Orm) JwtService {
	return &jwtService{
		orm,
	}
}
