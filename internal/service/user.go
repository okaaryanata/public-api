package service

import (
	"context"

	"github.com/okaaryanata/public-api/internal/domain"
	"github.com/okaaryanata/public-api/pkg/user"
)

type (
	UserService struct {
		userClient *user.UserClient
	}
)

func NewUserService(userClient *user.UserClient) *UserService {
	return &UserService{
		userClient: userClient,
	}
}

func (u *UserService) CreateUser(ctx context.Context, args *domain.CreateUserArgs) (*domain.UserResponse, error) {
	return u.userClient.CreateUser(ctx, args)
}

func (u *UserService) GetUserByID(ctx context.Context, userID uint) (*domain.UserResponse, error) {
	return u.userClient.GetUserByID(ctx, userID)
}
