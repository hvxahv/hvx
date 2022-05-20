package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/public/v1alpha1"
)

type Public interface {
	pb.PublicClient
}

type public struct {
	pb.PublicClient
}

func NewPublic(c *Client) Public {
	return &public{
		PublicClient: pb.NewPublicClient(c.conn),
	}
}
