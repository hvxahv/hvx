package account

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/account"
	"github.com/hvxahv/hvx/clientv1"
)

type Account interface {
	pb.AccountsClient
}

type account struct {
	pb.AccountsClient
}

func NewAccount(c *clientv1.Client) Account {
	return &account{
		AccountsClient: pb.NewAccountsClient(c.Conn),
	}
}
