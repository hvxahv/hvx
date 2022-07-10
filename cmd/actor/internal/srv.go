package internal

import (
	gw "github.com/hvxahv/hvx/APIs/grpc-gateway/v1alpha1/actor"
	a "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/actor"
	svc "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

type server struct {
	a.ActorServer
}

const (
	AccountsTable = "accounts"
	ActorsTable   = "actors"
)

const (
	serviceName = "actor"
)

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	a.RegisterActorServer(s, &server{})

	if err := gw.RegisterActorHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
