package container

import (
	"github.com/originbenntou/2929BE/account/domain/repository"
)

func (c Container) GetAccountRepository() repository.UserRepository {
	return repository.NewUserRepository()
}
