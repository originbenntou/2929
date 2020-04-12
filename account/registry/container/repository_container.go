package container

import (
	"github.com/originbenntou/2929BE/account/domain/repository"
	repo "github.com/originbenntou/2929BE/account/infrastructure/datastore/mysql"
	"github.com/originbenntou/2929BE/shared/mysql"
)

func (c Container) GetAccountRepository(db mysql.DBManager) repository.UserRepository {
	return repo.NewUserRepository(db)
}
