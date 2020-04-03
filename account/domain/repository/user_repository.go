package repository

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"errors"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/valueobject/grpc/request"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository interface {
	InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error)
}

type userRepository struct {
	//db adaptor.DBAdaptor
}

func NewUserRepository() UserRepository {
	return userRepository{}
}

func (r userRepository) InsertUser(ctx context.Context, req request.InsertUserRequest) (*model.User, error) {
	//db, err := adaptor.NewMysqlConnection(constant.Config)
	//if err != nil {
	//	return nil, err
	//}
	//defer func() {
	//	if err := db.Close(); err != nil {
	//
	//	}
	//}()

	db, err := sql.Open("mysql", "2929:2929@tcp(2929mysql:3306)/account")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	insert, err2 := db.Prepare("INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err2 != nil {
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
