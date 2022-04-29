package cli

import (
	act "github.com/hvxahv/hvx/api/grpc/proto/activity/v1alpha1"
	"google.golang.org/grpc"
)

type Activity interface {
	act.ActivityClient
	act.FollowClient
}

type activity struct {
	act.ActivityClient
	act.FollowClient
}

func NewActivity(conn *grpc.ClientConn) Activity {
	return &activity{
		ActivityClient: act.NewActivityClient(conn),
		FollowClient:   act.NewFollowClient(conn),
	}
}
