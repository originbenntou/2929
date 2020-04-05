package mysql

import (
	"context"
	"database/sql"
	"github.com/originbenntou/2929BE/account/constant"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/shared/adaptor"
	"time"
)

type userRepository struct {
	*sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r userRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

func (r userRepository) CreateUser(ctx context.Context, req *model.User) (uint64, error) {
	db, err := adaptor.NewMysqlConnection(constant.Config)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	insert, err := db.Prepare("INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := insert.Close(); err != nil {
			panic(err)
		}
	}()

	result, err := insert.Exec(req.Email, req.PassHash, req.Name, req.CompanyId, time.Now().Format("2006-1-2 15:04:05"), time.Now().Format("2006-1-2 15:04:05"))
	if err != nil {
		return 0, err
	}

	lid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lid), nil
}
