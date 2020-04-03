package registry

import (
	"google.golang.org/grpc"

	"github.com/originbenntou/2929BE/account/application/usecase"
	"github.com/originbenntou/2929BE/account/interfaces/handler"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
)

type HandlerRegistry struct {
	*grpc.Server
}

func NewHandlerRegistry(gs *grpc.Server) *HandlerRegistry {
	return &HandlerRegistry{gs}
}

func (h HandlerRegistry) RegisterAccountHandler(u usecase.UserUseCase) {
	pbAccount.RegisterUserServiceServer(h.Server, handler.NewUserHandler(u))
}
