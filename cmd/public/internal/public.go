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
	exist, err := clientv1.New(p.ctx, microsvc.AccountServiceName).IsExistAccount(name)
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	return exist.IsExist, nil
}

func (p *Public) GetActorByUsername(username string) (*actor.ActorData, error) {
	a, err := clientv1.New(p.ctx, microsvc.ActorServiceName).GetActorByUsername(username)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return a, err
}

func (p *Public) CreateAccount(username, mail, password, publicKey string) (*account.CreateResponse, error) {
	create, err := clientv1.New(p.ctx, microsvc.AccountServiceName).CreateAccount(username, mail, password, publicKey)
	if err != nil {
		return nil, err
	}
	return create, err
}
