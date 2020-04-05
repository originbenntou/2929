package registry

import (
	"database/sql"
	"github.com/originbenntou/2929BE/account/registry/container"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"google.golang.org/grpc"
)

type Registry interface {
	Register()
}

type registry struct {
	*grpc.Server
	*sql.DB
	container.Container //DI Container
}

func NewRegistry(gs *grpc.Server, db *sql.DB) Registry {
	return &registry{gs, db, container.Container{}}
}

func (r registry) Register() {
	pbAccount.RegisterUserServiceServer(r.Server,
		r.GetAccountService(
			r.GetAccountRepository(r.DB),
		),
	)

	//grpchealth?
}
