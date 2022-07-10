package saved

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/saved/v1alpha1"
	"github.com/hvxahv/hvx/clientv1"
)

type Saved interface {
	pb.SavedClient
}

type saved struct {
	pb.SavedClient
}

func NewSaved(c *clientv1.Client) Saved {
	return &saved{pb.NewSavedClient(c.conn)}
}
