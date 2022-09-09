package internal

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/auth"
	"github.com/hvxahv/hvx/errors"
	svc "github.com/hvxahv/hvx/microsvc"
)

type server struct {
	auth.AuthServer
}

const (
	serviceName = "auth"
)

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	auth.RegisterAuthServer(s, &server{})

	if err := auth.RegisterAuthHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
