package clientv1

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
)

type Device interface {
	pb.DevicesClient
}

type device struct {
	pb.DevicesClient
}

func NewDevice(c *Client) Device {
	return &device{
		DevicesClient: pb.NewDevicesClient(c.conn),
	}
}
