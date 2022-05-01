/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package device

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
)

const serverName = "device"

type device struct {
	pb.DevicesServer
	*Devices
}

// Run starts the server. It will block until the server is shutdown. If the server fails to start, it will return an error.
func Run() error {

	return nil
}
