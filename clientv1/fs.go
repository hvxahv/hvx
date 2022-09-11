package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/fs"
	"github.com/hvxahv/hvx/errors"
)

type FS interface {
	CreateFs(accountId, fileName, address string) (*pb.CreateResponse, error)
	GetFs(accountId, fileName string) (*pb.GetResponse, error)
	Delete(account, fileName string) (*pb.DeleteResponse, error)
}

func (svc *Svc) CreateFs(accountId, fileName, address string) (*pb.CreateResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.New(errors.ErrConnectDeviceRPCServer)
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

func (svc *Svc) GetFs(accountId, fileName string) (*pb.GetResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.New(errors.ErrConnectDeviceRPCServer)
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

func (svc *Svc) Delete(account, fileName string) (*pb.DeleteResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, errors.New(errors.ErrConnectDeviceRPCServer)
	}
	defer c.Close()

	f, err := pb.NewFsClient(c.Conn).Delete(svc.ctx, &pb.DeleteRequest{
		AccountId: account,
		FileName:  fileName,
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}
