package handler

import (
	"context"
	"github.com/originbenntou/2929BE/account/application/usecase"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
)

func NewUserHandler(uc usecase.UserUseCase) pbAccount.UserServiceServer {
	return &userHandler{uc}
}

type userHandler struct {
	usecase.UserUseCase
}

func (h userHandler) CreateUser(ctx context.Context, pbReq *pbAccount.CreateUserRequest) (*pbAccount.CreateUserResponse, error) {
	resp, err := h.UserUseCase.CreateUser(ctx, request.InsertUserRequest{
		Email:    pbReq.Email,
		Password: pbReq.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pbAccount.CreateUserResponse{
		User: &pbAccount.User{
			Id:           resp.Id,
			Email:        resp.Email,
			PasswordHash: resp.PasswordHash,
			Name:         resp.Name,
			CompanyId:    resp.CompanyId,
		},
	}, nil
}
