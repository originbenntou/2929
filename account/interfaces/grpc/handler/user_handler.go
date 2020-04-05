package handler

import (
	"context"
	"errors"
	"github.com/originbenntou/2929BE/account/application/usecase"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"golang.org/x/crypto/bcrypt"
)

func NewUserHandler(uc usecase.UserUseCase) pbAccount.UserServiceServer {
	return &userHandler{uc}
}

type userHandler struct {
	usecase.UserUseCase
}

func (h userHandler) CreateUser(ctx context.Context, pbReq *pbAccount.CreateUserRequest) (*pbAccount.CreateUserResponse, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pbReq.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("internal server request")
	}

	resp, err := h.UserUseCase.CreateUser(ctx, request.InsertUserRequest{
		Email:     pbReq.GetEmail(),
		PassHash:  passHash,
		Name:      pbReq.GetName(),
		CompanyId: pbReq.GetCompanyId(),
	})
	if err != nil {
		return nil, err
	}

	return &pbAccount.CreateUserResponse{
		User: &pbAccount.User{
			Id: resp.Id,
		},
	}, nil
}
