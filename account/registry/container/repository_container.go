package container

import (
	"database/sql"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/account/infrastructure/datastore/mysql"
)

func (c Container) GetAccountRepository(db *sql.DB) repository.UserRepository {
	return mysql.NewUserRepository(db)
}
