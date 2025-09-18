package services

import (
	"context"
	"fmt"
	"goravel_by_gin/app/http/requests"
	"goravel_by_gin/app/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/facades"
)

type PostService interface {
	FindAll(ctx context.Context) ([]models.Post, error)

	FindById(ctx context.Context, id uint) (*models.Post, error)

	Create(ctx context.Context, bodyData requests.PostRequest, userId uint, file filesystem.File) error
}

type postService struct {
	orm           orm.Orm
	uploadService UploadService
}

func NewPostService(orm orm.Orm, uploadService UploadService) PostService {
	return &postService{
		orm,
		uploadService,
	}
}

func (s *postService) FindAll(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	if err := s.orm.
		Query().
		With("Category").
		// With("User", func(query orm.Query) orm.Query {
		// 	return query.Where("is_active", true)
		// }).
		With("User").
		With("Tags").
		With("Comments").
		OrderByDesc("created_at").
		FindOrFail(&posts); err != nil {
		facades.Log().Errorf("failed to get all posts error message is: %s", err.Error())
		return nil, fmt.Errorf("failed to get all posts")
	}

	return posts, nil
}

func (s *postService) FindById(ctx context.Context, id uint) (*models.Post, error) {
	var post models.Post
	if err := s.orm.
		Query().
		With("Category").
		With("User").
		With("Tags").
		With("Comments").
		Where("id", id).
		FirstOrFail(&post); err != nil {
		facades.Log().Errorf("Failed to get post by #%v error message is: %v", id, err)
		return nil, fmt.Errorf("failed to get post by id")
	}

	return &post, nil
}

func (s *postService) Create(
	ctx context.Context,
	bodyData requests.PostRequest,
	userId uint,
	file filesystem.File,
) error {
	imageName := fmt.Sprintf("%s_%d", uuid.NewString(), time.Now().Unix())
	filename, err := file.StoreAs("storage/uploads", imageName)
	if err != nil {
		return err
	}

	post := &models.Post{
		Title:      bodyData.Title,
		Summary:    bodyData.Summary,
		Content:    bodyData.Content,
		IsActive:   bodyData.IsActive,
		Slug:       bodyData.Slug,
		ImageUrl:   facades.Storage().Url(filename),
		UserId:     userId,
		CategoryId: bodyData.CategoryId,
	}
	if err := s.orm.Query().Create(post); err != nil {
		if err := facades.Storage().Delete(filename); err != nil {
			return err
		}
		facades.Log().Errorf("failed to create new post: %v", err)
		return fmt.Errorf("failed to create new post")
	}

	if len(bodyData.Tags) > 0 {
		var tags []models.Tag
		for _, tagName := range bodyData.Tags {
			newTagName := strings.Split(tagName, ",")
			for _, tname := range newTagName {
				tagNameTrim := strings.Trim(tname, " ")
				var tag models.Tag
				if err := s.orm.Query().Where("name", tagNameTrim).FirstOrCreate(&tag, models.Tag{
					Name: tagNameTrim,
					Slug: tagNameTrim,
				}); err != nil {
					return err
				}
				tags = append(tags, tag)
			}
		}

		if err := s.orm.Query().Model(post).Association("Tags").Append(tags); err != nil {
			return err
		}

	}
	return nil
}
