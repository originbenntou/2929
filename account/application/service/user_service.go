package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserService interface {
	RegisterUser(context.Context, *pbAccount.RegisterUserRequest) (*pbAccount.RegisterUserResponse, error)
}

type userService struct {
	repository.UserRepository
	repository.CompanyRepository
}

func NewUserService(ur repository.UserRepository, cr repository.CompanyRepository) pbAccount.UserServiceServer {
	return &userService{ur, cr}
}

func (s userService) RegisterUser(ctx context.Context, pbReq *pbAccount.RegisterUserRequest) (*pbAccount.RegisterUserResponse, error) {
	user, err := s.FindUserByEmail(ctx, pbReq.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user != nil {
		return nil, status.Error(codes.AlreadyExists, errors.New("user already exist: "+pbReq.GetEmail()).Error())
	}

	company, err := s.FindCompanyById(ctx, pbReq.GetCompanyId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if company == nil {
		msg := fmt.Sprintf("company not found: %d", pbReq.GetCompanyId())
		return nil, status.Error(codes.InvalidArgument, errors.New(msg).Error())
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(pbReq.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	uid, err := s.CreateUser(ctx, &model.User{
		Email:     pbReq.GetEmail(),
		PassHash:  passHash,
		Name:      pbReq.GetName(),
		CompanyId: pbReq.GetCompanyId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbAccount.RegisterUserResponse{
		UserId: uid,
	}, nil
}
