package mysql

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/shared/logger"
	"github.com/originbenntou/2929BE/shared/mysql"
)

const InvalidID = 0

type userRepository struct {
	db mysql.DBManager
}

func NewUserRepository(db mysql.DBManager) repository.UserRepository {
	return &userRepository{db}
}

func (r userRepository) FindUserByEmail(ctx context.Context, email string) (u *model.User, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT * FROM user WHERE email = :email"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}

	// no match record is ok
	if !rows.Next() {
		return nil, nil
	}

	u = &model.User{}
	c := 0
	for rows.Next() {
		if err := rows.StructScan(u); err != nil {
			return nil, err
		}
		c++
	}

	if c > 1 {
		return nil, errors.New("found user more than 1 by: " + email)
	}

	// one match record
	return u, nil
}

func (r userRepository) CreateUser(ctx context.Context, req *model.User) (id uint64, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return InvalidID, err
	}

	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}

		if txErr := tx.CloseTransaction(err); txErr != nil {
			logger.Common.Error(txErr.Error())
		}
	}()

	q := "INSERT INTO user (email, password, name, company_id, created_at, updated_at) VALUES (:email, :password, :name, :company_id, :created_at, :updated_at)"

	result, err := tx.NamedExecContext(ctx, q, req)
	if err != nil {
		return 0, err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return InvalidID, err
	}

	if affect == 1 {
		return InvalidID, errors.New(fmt.Sprintf("total affected: %d", affect))
	}

	lid, err := result.LastInsertId()
	if err != nil {
		return InvalidID, err
	}

	return uint64(lid), nil
}
