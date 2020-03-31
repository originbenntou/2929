package repository

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
)

type UserRepository interface {
	InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error)
}
