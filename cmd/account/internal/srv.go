package internal

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/errors"
	svc "github.com/hvxahv/hvx/microsvc"
)

type server struct {
	account.AccountsServer
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

	account.RegisterAccountsServer(s, &server{})

	if err := account.RegisterAccountsHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
