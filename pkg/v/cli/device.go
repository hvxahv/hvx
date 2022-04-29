package cli

import (
	d "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
	"google.golang.org/grpc"
)

type Device interface {
	d.DevicesClient
}

type device struct {
	d.DevicesClient
}

func NewDevice(conn *grpc.ClientConn) Device {
	return &device{
		DevicesClient: d.NewDevicesClient(conn),
	}
}
