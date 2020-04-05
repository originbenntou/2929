package service

import (
	"context"
	"errors"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	RegisterUser(context.Context, *pbAccount.RegisterUserRequest) (*pbAccount.RegisterUserResponse, error)
}

type userService struct {
	repository.UserRepository
}

func NewUserService(repo repository.UserRepository) pbAccount.UserServiceServer {
	return &userService{repo}
}

func (s userService) RegisterUser(ctx context.Context, pbReq *pbAccount.RegisterUserRequest) (*pbAccount.RegisterUserResponse, error) {
	user, err := s.FindUserByEmail(ctx, pbReq.GetEmail())
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("User Already Exist: " + user.GetEmail())
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(pbReq.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("internal server request")
	}

	uid, err := s.CreateUser(ctx, &model.User{
		Email:     pbReq.GetEmail(),
		PassHash:  passHash,
		Name:      pbReq.GetName(),
		CompanyId: pbReq.GetCompanyId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return &pbAccount.RegisterUserResponse{
		UserId: uid,
	}, nil
}
