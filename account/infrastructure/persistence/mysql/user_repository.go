package mysql

import (
	"context"
	"github.com/originbenntou/2929BE/account/constant"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	"github.com/originbenntou/2929BE/shared/adaptor"
	"time"
)

type UserMySQL interface {
	InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error)
}

type userMySQL struct {
}

func NewUserMySQL() UserMySQL {
	return userMySQL{}
}

func (u userMySQL) InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error) {
	db, err := adaptor.NewMysqlConnection(constant.Config)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	insert, err := db.Prepare("INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := insert.Close(); err != nil {
			panic(err)
		}
	}()

	result, err := insert.Exec(req.Email, req.PassHash, "", "0", time.Now().Format("2006-1-2 15:04:05"), time.Now().Format("2006-1-2 15:04:05"))
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id: uint64(id),
	}, nil
}
