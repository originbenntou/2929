package mysql

import (
	"context"
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

func (r userRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

func (r userRepository) CreateUser(ctx context.Context, req *model.User) (uint64, error) {
	//defer func() {
	//	if err := r.db.Close(); err != nil {
	//		panic(err)
	//	}
	//}()

	q := "INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)"

	insert, err := r.db.Prepare(q)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := insert.Close(); err != nil {
			panic(err)
		}
	}()

	result, err := insert.Exec(req.Email, req.PassHash, req.Name, req.CompanyId, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return 0, err
	}

	lid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lid), nil
}
