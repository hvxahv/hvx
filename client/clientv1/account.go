package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/account/v1alpha1"
)

type Account interface {
	pb.AccountsClient
	pb.ActorClient
	pb.AuthClient
}

type account struct {
	pb.AccountsClient
	pb.ActorClient
	pb.AuthClient
}

func NewAccount(c *Client) Account {
	return &account{
		AccountsClient: pb.NewAccountsClient(c.conn),
		ActorClient:    pb.NewActorClient(c.conn),
		AuthClient:     pb.NewAuthClient(c.conn),
	}
}
