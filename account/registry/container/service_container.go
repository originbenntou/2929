package container

import (
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/account/domain/service"
)

func (c Container) GetAccountService(r repository.UserRepository) service.UserService {
	return service.NewUserService(r)
}
