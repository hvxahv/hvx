package internal

import (
	gw "github.com/hvxahv/hvx/APIs/grpc-gateway/v1alpha1/auth"
	auths "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/auth"
	svc "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

type server struct {
	auths.AuthServer
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

	auths.RegisterAuthServer(s, &server{})

	if err := gw.RegisterAuthHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
