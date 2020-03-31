package service

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/response"
)

type UserService interface {
	CreateUserService(ctx context.Context, req request.InsertUserRequest) (*response.InsertUserResponse, error)
}

type userService struct {
	repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return userService{r}
}

func (s userService) CreateUserService(ctx context.Context, req request.InsertUserRequest) (*response.InsertUserResponse, error) {
	resp, err := s.UserRepository.InsertUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &response.InsertUserResponse{
		Id:           resp.Id,
		Email:        resp.Email,
		PasswordHash: resp.Password,
		Name:         resp.Name,
		CompanyId:    resp.CompanyId,
	}, nil
}
