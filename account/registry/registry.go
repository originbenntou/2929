package registry

import (
	"github.com/originbenntou/2929BE/account/registry/container"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"github.com/originbenntou/2929BE/shared/mysql"
	"google.golang.org/grpc"
)

type Registry interface {
	Register()
}

type registry struct {
	*grpc.Server
	mysql.DBManager
	container.Container //DI Container
}

func NewRegistry(s *grpc.Server, db mysql.DBManager) Registry {
	return &registry{s, db, container.Container{}}
}

func (r registry) Register() {
	pbAccount.RegisterUserServiceServer(r.Server,
		r.GetAccountService(
			r.GetUserService(
				r.GetUserRepository(r.DBManager),
				r.GetCompanyRepository(r.DBManager),
			),
		),
	)

	//grpchealth?
}
