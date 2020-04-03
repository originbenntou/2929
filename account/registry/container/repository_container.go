package container

import (
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/account/infrastructure/persistence/mysql"
)

func (c Container) GetAccountRepository() repository.UserRepository {
	return mysql.NewUserRepository()
}
