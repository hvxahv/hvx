package internal

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/account/v1alpha1"
	svc "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

type server struct {
	pb.AccountsServer
	pb.ActorServer
	pb.AuthServer
}

const (
	AccountsTable = "accounts"
	ActorsTable   = "actors"
)
const (
	UsernameAlreadyExists = "THE_USERNAME_ALREADY_EXISTS"
)

const (
	serviceName = "account"
)

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	pb.RegisterAccountsServer(s, &server{})
	pb.RegisterActorServer(s, &server{})
	pb.RegisterAuthServer(s, &server{})

	if err := pb.RegisterActorHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := pb.RegisterAccountsHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
