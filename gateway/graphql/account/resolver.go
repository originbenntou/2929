package account

import (
	"context"
	"github.com/originbenntou/2929BE/gateway/grpc/client"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"google.golang.org/grpc"
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
		accountClient: pbAccount.NewUserServiceClient(client.GetGrpcConn(os.Getenv("ACCOUNT_ADDR"), xTraceID)),
	}
}

func xTraceID(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//traceID := support.GetTraceIDFromContext(ctx)
	//ctx = md.AddTraceIDToContext(ctx, traceID)
	return invoker(ctx, method, req, reply, cc, opts...)
}
