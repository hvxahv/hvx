package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/public"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Public interface {
	GetInstance() (*pb.GetInstanceResponse, error)
}

func (svc *Svc) GetInstance() (*pb.GetInstanceResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	instance, err := pb.NewPublicClient(c.Conn).GetInstance(svc.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return instance, nil
}
