package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"github.com/originbenntou/2929BE/shared/interceptor"
	"google.golang.org/grpc"
)

const port = ":50051"

type AccountService struct {
}

func (s *AccountService) CreateUser(ctx context.Context, req *pbAccount.CreateUserRequest) (*pbAccount.CreateUserResponse, error) {
	return nil, nil
}
func (s *AccountService) VerifyUser(ctx context.Context, req *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error) {
	return nil, nil
}
func (s *AccountService) FindUser(ctx context.Context, req *pbAccount.FindUserRequest) (*pbAccount.FindUserResponse, error) {
	return nil, nil
}

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
	pbAccount.RegisterUserServiceServer(srv, &AccountService{})

	go func() {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to create listener: %s",
				err)
		}
		log.Println("start server on port", port)
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
