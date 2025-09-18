package services

import (
	"context"
	"fmt"
	"goravel_by_gin/app/models"

	"github.com/goravel/framework/contracts/database/orm"
)

type TagService interface {
	FindAll(ctx context.Context) ([]models.Tag, error)
}

type tagService struct {
	orm orm.Orm
}

func NewTagService(orm orm.Orm) TagService {
	return &tagService{
		orm,
	}
}

func (s *tagService) FindAll(ctx context.Context) ([]models.Tag, error) {
	var tags []models.Tag
	if err := s.orm.WithContext(ctx).
		Query().
		OrderByDesc("created_at").
		FindOrFail(&tags); err != nil {
		return nil, fmt.Errorf("tag is not found")
	}
	return tags, nil
}
