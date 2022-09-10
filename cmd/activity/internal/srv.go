package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/errors"
	svc "github.com/hvxahv/hvx/microsvc"
)

type server struct {
	pb.ActivityServer
	pb.InboxServer
	pb.OutboxServer
	pb.FriendshipServer
}

const (
	serviceName = "activity"
)

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	pb.RegisterActivityServer(s, &server{})
	pb.RegisterInboxServer(s, &server{})
	pb.RegisterOutboxServer(s, &server{})
	//pb.RegisterFriendshipServer(s, &server{})

	if err := pb.RegisterActivityHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}
	if err := pb.RegisterInboxHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}
	if err := pb.RegisterOutboxHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}
	if err := pb.RegisterFriendshipHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
