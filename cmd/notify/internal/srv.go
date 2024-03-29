package internal

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/notify"
	"github.com/hvxahv/hvx/errors"
	v "github.com/hvxahv/hvx/microsvc"
)

type server struct {
	pb.NotifyServer
}

const serviceName = "notify"

func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha"),
		v.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterNotifyServer(s, &server{})
	if err := pb.RegisterNotifyHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
