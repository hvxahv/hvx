package internal

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/message"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

type server struct {
	pb.MessagesServer
}

const serviceName = "message"

func Run() error {
	s := microsvc.New(
		microsvc.WithServiceName(serviceName),
		microsvc.WithServiceVersion("v1alpha"),
		microsvc.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterMessagesServer(s, &server{})
	if err := pb.RegisterMessagesHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register message services: %v", err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
