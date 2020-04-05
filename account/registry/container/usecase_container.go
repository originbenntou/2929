package container

import (
	"github.com/originbenntou/2929BE/account/application/service"
	"github.com/originbenntou/2929BE/account/domain/repository"
)

func (c Container) GetAccountService(r repository.UserRepository) service.UserService {
	return service.NewUserService(r)
}
