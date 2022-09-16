package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/fs"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
)

type FS interface {
	CreateFs(accountId int64, fileName, address string) (*pb.CreateResponse, error)
	GetFs(accountId int64, fileName string) (*pb.GetResponse, error)
	Delete(accountId int64, fileName string) (*pb.DeleteResponse, error)
}

func (svc *Svc) CreateFs(accountId int64, fileName, address string) (*pb.CreateResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.NewFailedToConnect(microsvc.DeviceServiceName)
	}
	defer c.Close()
	f, err := pb.NewFsClient(c.Conn).Create(svc.ctx, &pb.CreateRequest{
		AccountId: accountId,
		FileName:  fileName,
		Address:   address,
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (svc *Svc) GetFs(accountId int64, fileName string) (*pb.GetResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.NewFailedToConnect(microsvc.DeviceServiceName)
	}
	defer c.Close()

	f, err := pb.NewFsClient(c.Conn).Get(svc.ctx, &pb.GetRequest{
		AccountId: accountId,
		FileName:  fileName,
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (svc *Svc) Delete(accountId int64, fileName string) (*pb.DeleteResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.NewFailedToConnect(microsvc.DeviceServiceName)
	}
	defer c.Close()

	f, err := pb.NewFsClient(c.Conn).Delete(svc.ctx, &pb.DeleteRequest{
		AccountId: accountId,
		FileName:  fileName,
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}
