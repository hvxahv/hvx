package public

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/public/v1alpha1"
	"github.com/hvxahv/hvx/clientv1"
)

type Public interface {
	pb.PublicClient
}

type public struct {
	pb.PublicClient
}

func NewPublic(c *clientv1.Client) Public {
	return &public{
		PublicClient: pb.NewPublicClient(c.conn),
	}
}
