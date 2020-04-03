package registry

import (
	"github.com/originbenntou/2929BE/account/registry/container"
	"google.golang.org/grpc"
)

type Registry interface {
	Register()
}

type registry struct {
	*grpc.Server
	container.Container
}

func NewRegistry(gs *grpc.Server) Registry {
	return &registry{gs, container.Container{}}
}

func (r registry) Register() {
	h := NewHandlerRegistry(r.Server)
	h.RegisterAccountHandler(
		r.GetAccountUsecase(
			r.GetAccountService(
				r.GetAccountRepository(),
			),
		),
	)

	//grpchealth?
}
