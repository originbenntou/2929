package mysql

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/originbenntou/2929BE/account/constant"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	"github.com/originbenntou/2929BE/shared/adaptor"
	"time"
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
	defer func() {
		if err := db.Close(); err != nil {

		}
	}()

	insert, err := db.Prepare("INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := insert.Close(); err != nil {

		}
	}()

	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("internal server request")
	}

	result, err := insert.Exec(req.Email, passHash, "", "0", time.Now().Format("2006-1-2 15:04:05"), time.Now().Format("2006-1-2 15:04:05"))
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
