package internal

import (
	"github.com/hvxahv/hvx/APIs/grpc/v1alpha1/account"
	"github.com/hvxahv/hvx/APIs/grpc/v1alpha1/actor"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

type Public struct {
	ctx context.Context
}

type AccountHandler interface {
	IsExist(name string) (bool, error)
	CreateAccount(username, mail, password, publicKey string) (*account.CreateResponse, error)
}

type ActorHandler interface {
	//GetActorByUsername(username string) (*actor.ActorData, error)
}

type AuthHandler interface {
	//Auth(username, password string) (*auth.VerifyResponse, error)
}

func NewPublic(ctx context.Context) *Public {
	return &Public{ctx: ctx}
}

func (p *Public) IsExist(name string) (bool, error) {
	c, err := clientv1.New(p.ctx,
		[]string{microsvc.GetGRPCServiceAddress("account")},
	)
	if err != nil {
		return false, err
	}
	defer c.Close()
	exist, err := account.NewAccountsClient(c.Conn).IsExist(p.ctx, &account.IsExistRequest{
		Username: name,
	})
	if err != nil {
		return false, err
	}
	return exist.IsExist, nil
}

// GetActorByUsername ...
func (p *Public) GetActorByUsername(username string) (*actor.ActorData, error) {
	c, err := clientv1.New(p.ctx,
		[]string{microsvc.GetGRPCServiceAddress("actor")},
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	a, err := actor.NewActorClient(c.Conn).GetActorByUsername(p.ctx, &actor.GetActorByUsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return a, err
}

func (p *Public) CreateAccount(username, mail, password, publicKey string) (*account.CreateResponse, error) {
	c, err := clientv1.New(p.ctx,
		[]string{microsvc.GetGRPCServiceAddress("account")},
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	create, err := account.NewAccountsClient(c.Conn).Create(p.ctx, &account.CreateRequest{
		Username:  username,
		Mail:      mail,
		Password:  password,
		PublicKey: publicKey,
	})
	if err != nil {
		return nil, err
	}
	return create, err
}
