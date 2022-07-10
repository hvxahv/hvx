package actor

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/actor"
	"github.com/hvxahv/hvx/clientv1"
)

type Actor interface {
	pb.ActorClient
}

type actor struct {
	pb.ActorClient
}

func NewActor(c *clientv1.Client) Actor {
	return &actor{
		ActorClient: pb.NewActorClient(c.Conn),
	}
}
