package repository

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	"github.com/originbenntou/2929BE/account/infrastructure/persistence/mysql"
)

type UserRepository interface {
	InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error)
}

type userRepository struct {
	mysql.UserMySQL
}

func NewUserRepository(mysql mysql.UserMySQL) UserRepository {
	return userRepository{mysql}
}

func (r userRepository) InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error) {
	u, err := r.UserMySQL.InsertUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return u, nil
}
