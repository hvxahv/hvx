/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */
package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"github.com/hvxahv/hvx/errors"
	svc "github.com/hvxahv/hvx/microsvc"
)

type server struct {
	pb.ChannelServer
	pb.AdministrativeServer
	pb.BroadcastServer
	pb.SubscriberServer
}

const (
	serviceName = "channel"
)

func Run() error {
	s := svc.New(
		svc.WithServiceName(serviceName),
		svc.WithServiceVersion("v1alpha"),
		svc.WithServiceID("serviceName"),
	).ListenerWithEndpoints()

	pb.RegisterAdministrativeServer(s, &server{})
	pb.RegisterBroadcastServer(s, &server{})
	pb.RegisterChannelServer(s, &server{})
	pb.RegisterSubscriberServer(s, &server{})

	if err := pb.RegisterAdministrativeHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}

	if err := pb.RegisterBroadcastHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}
	if err := pb.RegisterChannelHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}
	if err := pb.RegisterSubscriberHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register %s services: %v", serviceName, err)
	}
	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
