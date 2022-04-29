package cli

import (
	acct "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"google.golang.org/grpc"
)

type Account interface {
	acct.AccountsClient
	acct.ActorClient
	acct.AuthClient
}

type account struct {
	acct.AccountsClient
	acct.ActorClient
	acct.AuthClient
}

func NewAccount(conn *grpc.ClientConn) Account {
	return &account{
		AccountsClient: acct.NewAccountsClient(conn),
		ActorClient:    acct.NewActorClient(conn),
		AuthClient:     acct.NewAuthClient(conn),
	}
}
