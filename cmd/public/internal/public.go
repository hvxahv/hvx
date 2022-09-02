package internal

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/clientv1"
	microsvc "github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

type Public struct {
	ctx context.Context
}

type PublicHandler interface {
	// IsExist determines if the account name exists.
	IsExist(name string) (bool, error)

	// GetActorByUsername get the actor object data by account name.
	GetActorByUsername(username string) (*actor.ActorData, error)

	// CreateAccount public create account API, connects to the account service.
	CreateAccount(username, mail, password, publicKey string) (*account.CreateResponse, error)
}

func NewPublic(ctx context.Context) *Public {
	return &Public{ctx: ctx}
}

func (p *Public) IsExist(name string) (bool, error) {
	c, err := clientv1.New(p.ctx,
		microsvc.NewGRPCAddress("account").Get(),
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

func (p *Public) GetActorByUsername(username string) (*actor.ActorData, error) {
	c, err := clientv1.New(p.ctx, microsvc.NewGRPCAddress("actor").Get())
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
		microsvc.NewGRPCAddress("account").Get(),
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
