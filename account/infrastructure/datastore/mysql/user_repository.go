package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/shared/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct {
	db mysql.DBManager
}

func NewUserRepository(db mysql.DBManager) repository.UserRepository {
	return &userRepository{db}
}

func (r userRepository) FindUserByEmail(ctx context.Context, email string) (u *model.User, err error) {
	q := "SELECT * FROM user WHERE email = :email"

	var rows *sqlx.Rows
	rows, err = r.db.NamedQueryContext(ctx, q, map[string]interface{}{"email": email})
	if err != nil {
		return
	}

	// 0件はnilを返却
	if !rows.Next() {
		return
	}

	u = &model.User{}
	c := 0
	for rows.Next() {
		if err = rows.StructScan(u); err != nil {
			return
		}
		c++
	}

	// 2件以上はエラー
	if c > 1 {
		u = nil
		err = errors.New("found user more than 1 by:" + email)
		return
	}

	return
}

func (r userRepository) CreateUser(ctx context.Context, req *model.User) (id uint64, err error) {
	var tx mysql.TxManager
	tx, err = r.db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err = mysql.CloseTransaction(tx, err); err != nil {
			return
		}
	}()

	q := "INSERT INTO user (email, password, name, company_id, created_at, updated_at) VALUES (:email, :password, :name, :company_id, :created_at, :updated_at)"

	var result sql.Result
	result, err = tx.NamedExecContext(ctx, q, req)
	if err != nil {
		return
	}

	var affect int64
	affect, err = result.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		err = errors.New(fmt.Sprintf("total affected: %d", affect))
		return
	}

	var lid int64
	lid, err = result.LastInsertId()
	if err != nil {
		return
	}

	id = uint64(lid)

	return
}
