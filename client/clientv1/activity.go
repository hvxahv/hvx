package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/activity/v1alpha1"
)

type Activity interface {
	pb.ActivityClient
	pb.FollowClient
}

type activity struct {
	pb.ActivityClient
	pb.FollowClient
}

func NewActivity(c *Client) Activity {
	return &activity{
		ActivityClient: pb.NewActivityClient(c.conn),
		FollowClient:   pb.NewFollowClient(c.conn),
	}
}
