package main

import (
	"context"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/originbenntou/2929BE/account/constant"
	"github.com/originbenntou/2929BE/account/registry"
	"github.com/originbenntou/2929BE/shared/interceptor"
	"github.com/originbenntou/2929BE/shared/logger"
	"github.com/originbenntou/2929BE/shared/mysql"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			interceptor.XTraceID(),
			grpc_zap.UnaryServerInterceptor(logger.Logger),
			interceptor.Logging(),
		)),
	)

	conn, err := mysql.NewDBConnection(constant.Config)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	// DB操作をラップ
	m := mysql.NewDBManager(conn)

	registry.NewRegistry(srv, m).Register()
	reflection.Register(srv)

	go func() {
		listener, err := net.Listen("tcp", constant.Port)
		if err != nil {
			log.Fatalf("failed to create listener: %s", err)
		}
		log.Println("start server on port", constant.Port)
		if err := srv.Serve(listener); err != nil {
			log.Println("failed to exit serve: ", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM)
	<-sigint
	log.Println("received a signal of graceful shutdown")

	stopped := make(chan struct{})
	go func() {
		srv.GracefulStop()
		close(stopped)
	}()

	ctx, cancel := context.WithTimeout(
		context.Background(), 1*time.Minute)

	select {
	case <-ctx.Done():
		srv.Stop()
	case <-stopped:
		cancel()
	}

	log.Println("completed graceful shutdown")
}
