package services

import (
	"context"
	"fmt"
	"goravel_by_gin/app/enums"
	"goravel_by_gin/app/http/requests"
	"goravel_by_gin/app/models"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type AuthService interface {
	Register(ctx context.Context, bodyData *requests.AuthRequest) (*models.User, error)

	Login(ctx context.Context, bodyData *requests.AuthLoginRequest) (*models.User, error)

	FindUserById(ctx context.Context, id uint) (*models.User, error)

	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type authService struct {
	// Dependency
	orm orm.Orm
}

func NewAuthService(orm orm.Orm) AuthService {
	return &authService{
		// Inject dependency
		orm,
	}
}

func (s *authService) Register(ctx context.Context, bodyData *requests.AuthRequest) (*models.User, error) {
	hashed, err := facades.Hash().Make(bodyData.Password)
	if err != nil {
		facades.Log().Errorf("failed to hash user password error message: %v", err)
		return nil, fmt.Errorf("failed to hash password")
	}

	user := &models.User{
		Name:     bodyData.Name,
		Email:    bodyData.Email,
		Password: hashed,
		IsActive: true,
		Role:     enums.RoleGuest,
	}
	if err := s.orm.Query().Create(user); err != nil {
		facades.Log().Errorf("failed to register user error message: %v", err)
		return nil, fmt.Errorf("failed to register user")
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, bodyData *requests.AuthLoginRequest) (*models.User, error) {
	user, err := s.FindUserByEmail(ctx, bodyData.Email)
	if err != nil {
		return nil, err
	}

	if !facades.Hash().Check(bodyData.Password, user.Password) {
		return nil, fmt.Errorf("email or password is invalid")
	}

	return user, nil
}

func (s *authService) FindUserById(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := s.orm.Query().Where("id", id).FirstOrFail(&user); err != nil {
		facades.Log().Errorf("failed to get current user error message: %v", err)
		return nil, fmt.Errorf("failed to get current user")
	}

	return &user, nil
}

func (s *authService) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := s.orm.Query().Where("email", email).FirstOrFail(&user); err != nil {
		return nil, fmt.Errorf("failed to get user by email")
	}
	return &user, nil
}
