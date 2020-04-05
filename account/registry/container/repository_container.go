package container

import (
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/account/infrastructure/persistence/mysql"
)

func (c Container) GetAccountRepository(mysql mysql.UserMySQL) repository.UserRepository {
	return repository.NewUserRepository(mysql)
}
