package usecase

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/repository"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, req *pbAccount.CreateUserRequest) (*pbAccount.CreateUserResponse, error)
	VerifyUser(ctx context.Context, req *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error)
	FindUser(ctx context.Context, req *pbAccount.FindUserRequest) (*pbAccount.FindUserResponse, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) CreateUser(ctx context.Context, req *pbAccount.CreateUserRequest) (*pbAccount.CreateUserResponse, error) {
	return nil, nil
}

func (u *userUseCase) VerifyUser(ctx context.Context, req *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error) {
	return nil, nil
}

func (u *userUseCase) FindUser(ctx context.Context, req *pbAccount.FindUserRequest) (*pbAccount.FindUserResponse, error) {
	return nil, nil
}
