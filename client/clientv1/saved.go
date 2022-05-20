package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/saved/v1alpha1"
)

type Saved interface {
	pb.SavedClient
}

type saved struct {
	pb.SavedClient
}

func NewSaved(c *Client) Saved {
	return &saved{pb.NewSavedClient(c.conn)}
}
