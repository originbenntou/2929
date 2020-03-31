package usecase

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/service"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/response"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, req request.InsertUserRequest) (*response.InsertUserResponse, error)
}

type userUseCase struct {
	service.UserService
}

func NewUserUseCase(s service.UserService) UserUseCase {
	return userUseCase{s}
}

func (u userUseCase) CreateUser(ctx context.Context, req request.InsertUserRequest) (*response.InsertUserResponse, error) {
	return u.UserService.CreateUserService(ctx, req)
}
