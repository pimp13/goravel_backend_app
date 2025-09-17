package services

import (
	"context"
	"fmt"
	"goravel_by_gin/app/http/requests"
	"goravel_by_gin/app/models"
	"time"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type PostService interface {
	FindAll(ctx context.Context) ([]models.Post, error)

	FindById(ctx context.Context, id uint) (*models.Post, error)

	Create(ctx context.Context, bodyData requests.PostRequest) error
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
		With("User").
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
		Where("id", id).
		FirstOrFail(&post); err != nil {
		facades.Log().Errorf("Failed to get post by #%v error message is: %v", id, err)
		return nil, fmt.Errorf("failed to get post by id")
	}

	return &post, nil
}

func (s *postService) Create(ctx context.Context, bodyData requests.PostRequest) error {
	imageName := fmt.Sprintf("%s_%d", uuid.NewString(), time.Now().Unix())
	filename, err := bodyData.Image.StoreAs("storage/uploads", imageName)
	if err != nil {
		return err
	}
	if err := s.orm.Query().Create(&models.Post{
		Title:      bodyData.Title,
		Summary:    bodyData.Summary,
		Content:    bodyData.Content,
		IsActive:   bodyData.IsActive,
		Slug:       bodyData.Slug,
		ImageUrl:   facades.Storage().Url(filename),
		UserId:     bodyData.UserId,
		CategoryId: bodyData.CategoryId,
	}); err != nil {
		return fmt.Errorf("failed to create a new category")
	}
	return nil
}
