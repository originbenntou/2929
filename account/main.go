package main

import (
	"context"
	"github.com/originbenntou/2929BE/account/constant"
	"github.com/originbenntou/2929BE/account/registry"
	"github.com/originbenntou/2929BE/shared/adaptor"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/originbenntou/2929BE/shared/interceptor"
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
			interceptor.Logging(),
		)),
	)

	conn, err := adaptor.NewMysqlConnection(constant.Config)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	registry.NewRegistry(srv, conn).Register()
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
