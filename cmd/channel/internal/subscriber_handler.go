package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"golang.org/x/net/context"
)

// Subscriber ...

func (s *server) AddSubscriber(ctx context.Context, in *pb.AddSubscriberRequest) (*pb.AddSubscriberResponse, error) {
	return &pb.AddSubscriberResponse{}, nil
}

func (s *server) RemoveSubscriber(ctx context.Context, in *pb.RemoveSubscriberRequest) (*pb.RemoveSubscriberResponse, error) {
	return &pb.RemoveSubscriberResponse{}, nil
}

func (s *server) GetSubscribers(ctx context.Context, in *pb.GetSubscribersRequest) (*pb.GetSubscribersResponse, error) {
	return &pb.GetSubscribersResponse{}, nil
}

func (s *server) Subscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	return &pb.SubscriptionResponse{}, nil
}

func (s *server) Unsubscribe(ctx context.Context, in *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	return &pb.UnsubscribeResponse{}, nil
}
