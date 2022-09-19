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
	// IsExist determines if the actor name exists.
	IsExist(name string) (*actor.IsExistResponse, error)

	// GetActorByUsername get the actor object data by account name.
	GetActorByUsername(username string) (*actor.ActorData, error)

	// CreateAccount public create account API, connects to the account service.
	CreateAccount(username, mail, password, publicKey string) (*account.CreateResponse, error)
}

func NewPublic(ctx context.Context) *Public {
	return &Public{ctx: ctx}
}

func (p *Public) IsExist(name string) (*actor.IsExistResponse, error) {
	exist, err := clientv1.New(p.ctx, microsvc.ActorServiceName).IsExistActor(name)
	if err != nil {
		return nil, err
	}
	return &actor.IsExistResponse{
		IsExist:   exist.IsExist,
		ActorType: exist.ActorType,
	}, nil
}

func (p *Public) GetActorByUsername(username string) (*actor.ActorData, error) {
	a, err := clientv1.New(p.ctx, microsvc.ActorServiceName).GetActorByUsername(username)
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
