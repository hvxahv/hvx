package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/errors"
	svc "github.com/hvxahv/hvx/microsvc"
)

type server struct {
	pb.ActorServer
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

	pb.RegisterActorServer(s, &server{})

	if err := pb.RegisterActorHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
