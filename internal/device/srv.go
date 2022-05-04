/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package device

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
	"github.com/hvxahv/hvx/pkg/v"
)

const (
	serviceName = "device"
)

type server struct {
	pb.DevicesServer
	*Devices
}

// Run starts the server. It will block until the server is shutdown. If the server fails to start, it will return an error.
func Run() error {
	s := v.New(
		v.WithServiceName(serviceName),
		v.WithServiceVersion("v1alpha"),
		v.WithServiceID(uuid.New().String()),
	).NewServer()

	pb.RegisterDevicesServer(s, &server{})
	if err := s.Run(); err != nil {
		return err
	}

	return nil
}
