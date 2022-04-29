package cli

import (
	pub "github.com/hvxahv/hvx/api/grpc/proto/public/v1alpha1"
	"google.golang.org/grpc"
)

type Public interface {
	pub.PublicClient
}

type public struct {
	pub.PublicClient
}

func NewPublic(conn *grpc.ClientConn) Public {
	return &public{
		PublicClient: pub.NewPublicClient(conn),
	}
}
