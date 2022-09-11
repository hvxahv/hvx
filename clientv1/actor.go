package clientv1

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
)

type Actor interface {
	GetActor(actorId string) (*actor.GetResponse, error)
	GetActorByAddress(inbox string) (*actor.ActorData, error)
}

func (svc *Svc) GetActor(actorId string) (*actor.GetResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	a, err := actor.NewActorClient(c.Conn).Get(svc.ctx, &actor.GetRequest{
		ActorId: actorId,
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (svc *Svc) GetActorByAddress(inbox string) (*actor.ActorData, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	a, err := actor.NewActorClient(c.Conn).GetActorByAddress(svc.ctx, &actor.GetActorByAddressRequest{
		Address: inbox,
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}
