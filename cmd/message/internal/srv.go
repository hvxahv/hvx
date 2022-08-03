package internal

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/message"
	v "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

type server struct {
	pb.MessagesServer
}

const serviceName = "message"

func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha"),
		v.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterMessagesServer(s, &server{})
	if err := pb.RegisterMessagesHandlerClient(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
