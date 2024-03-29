package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/device"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
)

type Devices interface {
	DeleteDevices(accountId int64) (*pb.DeleteDevicesResponse, error)
	AddDevice(accountId int64, userAgent string) (*pb.CreateResponse, error)
	IsExistDevice(deviceId int64) (*pb.IsExistResponse, error)
}

func (svc *Svc) DeleteDevices(accountId int64) (*pb.DeleteDevicesResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.NewFailedToConnect(microsvc.DeviceServiceName)
	}
	defer c.Close()
	d, err := pb.NewDevicesClient(c.Conn).DeleteDevices(svc.ctx, &pb.DeleteDevicesRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (svc *Svc) AddDevice(accountId int64, userAgent string) (*pb.CreateResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.NewFailedToConnect(microsvc.DeviceServiceName)
	}
	defer c.Close()
	d, err := pb.NewDevicesClient(c.Conn).Create(svc.ctx, &pb.CreateRequest{
		AccountId: accountId,
		Ua:        userAgent,
	})
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (svc *Svc) IsExistDevice(deviceId int64) (*pb.IsExistResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.NewFailedToConnect(microsvc.DeviceServiceName)
	}
	defer c.Close()
	devices, err := pb.NewDevicesClient(c.Conn).IsExist(svc.ctx, &pb.IsExistRequest{
		Id: deviceId,
	})
	if err != nil {
		return nil, err
	}
	return devices, nil
}
