/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package device

import (
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
)

func GetDeviceClient() (v1alpha1.DevicesClient, error) {
	conn, err := grpc.Dial(microservices.NewService("device").GetAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return v1alpha1.NewDevicesClient(conn), nil
}
