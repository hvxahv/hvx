package activity

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/activity"
	"github.com/hvxahv/hvx/clientv1"
)

type Activity interface {
	pb.ActivityClient
	pb.FollowClient
}

type activity struct {
	pb.ActivityClient
	pb.FollowClient
}

func NewActivity(c *clientv1.Client) Activity {
	return &activity{
		ActivityClient: pb.NewActivityClient(c.conn),
		FollowClient:   pb.NewFollowClient(c.conn),
	}
}
