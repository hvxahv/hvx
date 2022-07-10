/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/google/uuid"
	gw "github.com/hvxahv/hvx/APIs/grpc-gateway/v1alpha1/device"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/device"
	v "github.com/hvxahv/hvx/microsvc"
	"github.com/pkg/errors"
)

const (
	serviceName = "device"
)

type server struct {
	pb.DevicesServer
}

func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha1"),
		v.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterDevicesServer(s, &server{})
	if err := s.Run(); err != nil {
		return err
	}
	if err := gw.RegisterDevicesHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Errorf("Failed to register %s services: %v", serviceName, err)
	}

	return nil
}
