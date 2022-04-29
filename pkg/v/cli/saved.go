package cli

import (
	save "github.com/hvxahv/hvx/api/grpc/proto/saved/v1alpha1"
	"google.golang.org/grpc"
)

type Saved interface {
	save.SavedClient
}

type saved struct {
	save.SavedClient
}

func NewSaved(conn *grpc.ClientConn) Saved {
	return &saved{save.NewSavedClient(conn)}
}
