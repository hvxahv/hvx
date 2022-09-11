package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/actor"
)

type Actor interface {
	IsExistActor(preferredUsername string) (*pb.IsExistResponse, error)
	GetActor(actorId string) (*pb.GetResponse, error)
	GetActorByUsername(username string) (*pb.ActorData, error)
	GetActorByAddress(inbox string) (*pb.ActorData, error)
	CreateActor(preferredUsername, publicKey, actorType string) (*pb.CreateResponse, error)
	DeleteActor(actorId string) (*pb.DeleteResponse, error)
}

func (svc *Svc) IsExistActor(preferredUsername string) (*pb.IsExistResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	i, err := pb.NewActorClient(c.Conn).IsExist(svc.ctx, &pb.IsExistRequest{
		PreferredUsername: preferredUsername,
	})
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (svc *Svc) GetActor(actorId string) (*pb.GetResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	a, err := pb.NewActorClient(c.Conn).Get(svc.ctx, &pb.GetRequest{
		ActorId: actorId,
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (svc *Svc) GetActorByUsername(username string) (*pb.ActorData, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	d, err := pb.NewActorClient(c.Conn).GetActorByUsername(svc.ctx, &pb.GetActorByUsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (svc *Svc) GetActorByAddress(inbox string) (*pb.ActorData, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	a, err := pb.NewActorClient(c.Conn).GetActorByAddress(svc.ctx, &pb.GetActorByAddressRequest{
		Address: inbox,
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (svc *Svc) CreateActor(preferredUsername, publicKey, actorType string) (*pb.CreateResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	create, err := pb.NewActorClient(c.Conn).Create(svc.ctx, &pb.CreateRequest{
		PreferredUsername: preferredUsername,
		PublicKey:         publicKey,
		ActorType:         actorType,
	})
	if err != nil {
		return nil, err
	}

	return create, nil
}

func (svc *Svc) DeleteActor(actorId string) (*pb.DeleteResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	d, err := pb.NewActorClient(c.Conn).Delete(svc.ctx, &pb.DeleteRequest{
		Id: actorId,
	})
	if err != nil {
		return nil, err
	}

	return d, nil
}
