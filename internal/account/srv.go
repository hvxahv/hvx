package account

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	v "github.com/hvxahv/hvx/pkg/microsvc"
	"github.com/pkg/errors"
)

type server struct {
	pb.AccountsServer
	pb.ActorServer
	pb.AuthServer
	*Accounts
	*Actors
}

const (
	serviceName = "account"
)

func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha"),
		v.WithServiceID(uuid.New().String()),
	).NewServer()

	pb.RegisterAccountsServer(s, &server{})
	pb.RegisterActorServer(s, &server{})
	pb.RegisterAuthServer(s, &server{})

	if err := pb.RegisterActorHandler(s.GetCtx(), s.GetMux(), s.GetConn()); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := pb.RegisterAccountsHandler(s.GetCtx(), s.GetMux(), s.GetConn()); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
