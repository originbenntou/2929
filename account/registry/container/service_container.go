package container

import (
	"github.com/originbenntou/2929BE/account/application/service"
	"github.com/originbenntou/2929BE/account/domain/repository"
)

func (c Container) GetUserService(ur repository.UserRepository, cr repository.CompanyRepository) service.UserService {
	return service.NewUserService(ur, cr)
}
