/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/device"
	"github.com/hvxahv/hvx/errors"
	v "github.com/hvxahv/hvx/microsvc"
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
		v.WithServiceVersion("v1alpha"),
		v.WithServiceID(uuid.New().String()),
	).ListenerWithEndpoints()

	pb.RegisterDevicesServer(s, &server{})
	if err := pb.RegisterDevicesHandler(s.Ctx, s.Mux, s.Conn); err != nil {
		return errors.Newf("Failed to register public services: %v", err)
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
