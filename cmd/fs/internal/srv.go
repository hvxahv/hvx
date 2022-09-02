package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/fs"
	svc "github.com/hvxahv/hvx/microsvc"
)

type server struct {
	pb.FsServer
}

const (
	serviceName = "fs"
)

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	pb.RegisterFsServer(s, &server{})

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
