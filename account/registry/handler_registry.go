package registry

import (
	"github.com/originbenntou/2929BE/account/interfaces/grpc/handler"
	"google.golang.org/grpc"

	"github.com/originbenntou/2929BE/account/application/usecase"
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
