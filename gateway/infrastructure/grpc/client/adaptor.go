package client

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/originbenntou/2929BE/shared/logger"
	"google.golang.org/grpc"
)

func GetGrpcConn(target string, interceptors ...grpc.UnaryClientInterceptor) *grpc.ClientConn {
	chain := grpc_middleware.ChainUnaryClient(interceptors...)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithUnaryInterceptor(chain))
	if err != nil {
		logger.Common.Error(err.Error())
	}
	return conn
}
