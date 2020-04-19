package account

import (
	"github.com/originbenntou/2929BE/gateway/infrastructure/grpc/client"
	"github.com/originbenntou/2929BE/gateway/interfaces/interceptor"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"os"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	accountClient pbAccount.UserServiceClient
}

func NewAccountResolver() *Resolver {
	return &Resolver{
		accountClient: pbAccount.NewUserServiceClient(client.GetGrpcConn(os.Getenv("ACCOUNT_ADDR"), interceptor.XTraceID)),
	}
}
