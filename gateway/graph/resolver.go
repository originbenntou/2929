package graph

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"google.golang.org/grpc"
	"log"
	"os"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userClient pbAccount.UserServiceClient
	//trendClient pbTrend.suggestServiceClient
}

func NewGraphQLResolver() *Resolver {
	return &Resolver{
		userClient: pbAccount.NewUserServiceClient(getGRPCConn(os.Getenv("USER_SERVICE_ADDR"), xTraceID)),
	}
}

func getGRPCConn(target string, interceptors ...grpc.UnaryClientInterceptor) *grpc.ClientConn {
	// インタセプタを使ってRPC共通の処理の追加
	chain := grpc_middleware.ChainUnaryClient(interceptors...)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithUnaryInterceptor(chain))
	if err != nil {
		log.Fatalf("failed to dial: %s", err)
	}
	return conn
}

func xTraceID(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	//traceID := support.GetTraceIDFromContext(ctx)
	//ctx = md.AddTraceIDToContext(ctx, traceID)
	return invoker(ctx, method, req, reply, cc, opts...)
}
