package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
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
	VerifyUser(context.Context, *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error)
}

type userService struct {
	repository.UserRepository
	repository.CompanyRepository
	repository.SessionRepository
	repository.PlanRepository
}

func NewUserService(
	ur repository.UserRepository,
	cr repository.CompanyRepository,
	sr repository.SessionRepository,
	pr repository.PlanRepository,
) UserService {
	return &userService{
		ur,
		cr,
		sr,
		pr,
	}
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

func (s userService) VerifyUser(ctx context.Context, pbReq *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error) {
	user, err := s.FindUserByEmail(ctx, pbReq.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, errors.New("user is not found: "+pbReq.GetEmail()).Error())
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(pbReq.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// forbidden login over limit by plan
	isOver, err := s.isLoginOverCapacityOfCompanyPlan(ctx, user.CompanyId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if isOver {
		return nil, status.Error(codes.Unauthenticated, errors.New("forbidden login over limit by plan").Error())
	}

	// if double login, continue to use old token
	// if new login, generate new token
	token, err := s.getCorrectTokenForUser(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbAccount.VerifyUserResponse{
		Token: token,
		User: &pbAccount.User{
			Id:        user.Id,
			Email:     user.Email,
			Name:      user.Name,
			CompanyId: user.CompanyId,
		},
	}, nil
}

func (s userService) isLoginOverCapacityOfCompanyPlan(ctx context.Context, cid uint64) (bool, error) {
	count, err := s.CountValidSessionByCompanyId(ctx, cid)
	if err != nil {
		return false, err
	}

	capacity, err := s.FindCapacityByCompanyId(ctx, cid)
	if err != nil {
		return false, err
	}

	return count > capacity, nil
}

func (s userService) getCorrectTokenForUser(ctx context.Context, user *model.User) (string, error) {
	oldToken, err := s.FindExistTokenByUserId(ctx, user.Id)
	if err != nil {
		return "", err
	}

	// if token already exists, it is double login
	if oldToken != "" {
		// update existing session to extend expire
		if err := s.UpdateSession(ctx, user.Id); err != nil {
			return "", err
		}

		// continue to use old token
		return oldToken, nil
	}

	// if new login
	newToken := uuid.New().String()
	// create new session
	if err = s.CreateSession(ctx, &model.Session{
		Token:     newToken,
		UserId:    user.Id,
		CompanyId: user.CompanyId,
	}); err != nil {
		return "", err
	}

	return newToken, nil
}
