package container

import (
	"github.com/originbenntou/2929BE/account/application/service"
	"github.com/originbenntou/2929BE/account/interfaces/handler"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
)

func (c Container) GetAccountService(su service.UserService) pbAccount.UserServiceServer {
	return handler.NewAccountHandler(su)
}
