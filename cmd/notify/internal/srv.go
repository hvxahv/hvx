package internal

import (
	"github.com/google/uuid"
	gw "github.com/hvxahv/hvx/APIs/grpc-gateway/v1alpha1/notify"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/notify"
	v "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
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
	if err := gw.RegisterNotifyHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
