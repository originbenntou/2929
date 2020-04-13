package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/shared/mysql"
)

type userRepository struct {
	db mysql.DBManager
}

func NewUserRepository(db mysql.DBManager) repository.UserRepository {
	return &userRepository{db}
}

func (r userRepository) FindUserByEmail(ctx context.Context, email string) (u *model.User, err error) {
	var rows *sql.Rows
	q := "SELECT * FROM user WHERE email = ?"
	rows, err = r.db.QueryContext(ctx, q, email)
	if err != nil {
		return
	}

	u = &model.User{}
	for rows.Next() {
		if err = rows.Scan(&u.Id, &u.Email, &u.PassHash, &u.Name, &u.CompanyId, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return
		}
	}

	// ユーザーが存在しない場合はIdゼロのUser構造体を返す
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

	q := "INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)"

	var stmt *sql.Stmt
	stmt, err = tx.PrepareContext(ctx, q)
	if err != nil {
		return
	}

	defer func() {
		if stmtErr := stmt.Close(); stmtErr != nil {
			err = stmtErr
			return
		}
	}()

	var result sql.Result
	result, err = stmt.ExecContext(ctx, req.Email, req.PassHash, req.Name, req.CompanyId, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return
	}

	var affect int64
	affect, err = result.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		err = errors.New(fmt.Sprintf("total affected: %d ", affect))
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
