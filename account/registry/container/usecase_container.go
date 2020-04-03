package container

import (
	"github.com/originbenntou/2929BE/account/application/usecase"
	"github.com/originbenntou/2929BE/account/domain/service"
)

func (c Container) GetAccountUsecase(s service.UserService) usecase.UserUseCase {
	return usecase.NewUserUseCase(s)
}
