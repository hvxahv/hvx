/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/google/uuid"
	gw "github.com/hvxahv/hvx/APIs/grpc-gateway/v1alpha1/channel"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/channel"
	v "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

const (
	serviceName = "channel"
)

type server struct {
	pb.AdministrativeServer
	pb.BroadcastServer
	pb.ChannelServer
	pb.SubscriberServer
}

func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha1"),
		v.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterAdministrativeServer(s, &server{})
	pb.RegisterBroadcastServer(s, &server{})
	pb.RegisterChannelServer(s, &server{})
	pb.RegisterSubscriberServer(s, &server{})

	if err := s.Run(); err != nil {
		return err
	}

	if err := gw.RegisterAdministrativeHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}
	if err := gw.RegisterBroadcastHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}
	if err := gw.RegisterChannelHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}
	if err := gw.RegisterSubscriberHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	return nil
}
