package mysql

import (
	"context"
	"github.com/originbenntou/2929BE/account/constant"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	"github.com/originbenntou/2929BE/shared/adaptor"
)

type UserMySQL interface {
	InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error)
}

type userRepository struct {
	//db adaptor.DBAdaptor
}

func NewUserRepository() repository.UserRepository {
	return userRepository{}
}

func (r userRepository) InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error) {
	db, err := adaptor.NewMysqlConnection(constant.Config)
	if err != nil {
		return nil, err
	}
	db.Query()
	return nil, nil
}
