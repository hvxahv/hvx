package auth

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/auth"
	"github.com/hvxahv/hvx/clientv1"
)

type Auth interface {
	pb.AuthClient
}

type auth struct {
	pb.AuthClient
}

func NewAuth(c *clientv1.Client) Auth {
	return &auth{
		AuthClient: pb.NewAuthClient(c.Conn),
	}
}
