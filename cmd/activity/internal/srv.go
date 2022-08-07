package internal

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	svc "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

const serviceName = "activity"

type server struct {
	activity.ActivityServer
}

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	activity.RegisterActivityServer(s, &server{})

	if err := activity.RegisterActivityHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
