package device

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/device/v1alpha1"
	"github.com/hvxahv/hvx/clientv1"
)

type Device interface {
	pb.DevicesClient
}

type device struct {
	pb.DevicesClient
}

func NewDevice(c *clientv1.Client) Device {
	return &device{
		DevicesClient: pb.NewDevicesClient(c.conn),
	}
}
