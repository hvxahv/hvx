package internal

import (
	"context"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/channel"
)

// Administrative ...

func (s *server) IsAdministrator(ctx context.Context, in *pb.IsAdministratorRequest) (*pb.IsAdministratorResponse, error) {
	return &pb.IsAdministratorResponse{}, nil
}

func (s *server) IsOwner(ctx context.Context, in *pb.IsOwnerRequest) (*pb.IsOwnerResponse, error) {
	return &pb.IsOwnerResponse{}, nil
}

func (s *server) AddAdministrator(ctx context.Context, in *pb.AddAdministratorRequest) (*pb.AddAdministratorResponse, error) {
	return &pb.AddAdministratorResponse{}, nil
}

func (s *server) RemoveAdministrator(ctx context.Context, in *pb.RemoveAdministratorRequest) (*pb.RemoveAdministratorResponse, error) {
	return &pb.RemoveAdministratorResponse{}, nil
}

func (s *server) GetAdministrators(ctx context.Context, in *pb.GetAdministratorsRequest) (*pb.GetAdministratorsResponse, error) {
	return &pb.GetAdministratorsResponse{}, nil
}

func (s *server) ExitAdministrator(ctx context.Context, in *pb.ExitAdministratorRequest) (*pb.ExitAdministratorResponse, error) {
	return &pb.ExitAdministratorResponse{}, nil
}

// Broadcast ...

func (s *server) CreateBroadcast(ctx context.Context, in *pb.CreateBroadcastRequest) (*pb.CreateBroadcastResponse, error) {
	return &pb.CreateBroadcastResponse{}, nil
}

func (s *server) GetBroadcasts(ctx context.Context, in *pb.GetBroadcastsRequest) (*pb.GetBroadcastsResponse, error) {
	return &pb.GetBroadcastsResponse{}, nil

}

func (s *server) DeleteBroadcast(ctx context.Context, in *pb.DeleteBroadcastRequest) (*pb.DeleteBroadcastResponse, error) {
	return &pb.DeleteBroadcastResponse{}, nil
}

// Channel ...

func (s *server) CreateChannel(ctx context.Context, in *pb.CreateChannelRequest) (*pb.CreateChannelResponse, error) {
	return &pb.CreateChannelResponse{}, nil
}

func (s *server) GetChannels(ctx context.Context, in *pb.GetChannelsRequest) (*pb.GetChannelsResponse, error) {
	return &pb.GetChannelsResponse{}, nil
}

func (s *server) DeleteChannel(ctx context.Context, in *pb.DeleteChannelRequest) (*pb.DeleteChannelResponse, error) {
	return &pb.DeleteChannelResponse{}, nil
}

func (s *server) DeleteChannels(ctx context.Context, in *pb.DeleteChannelsRequest) (*pb.DeleteChannelsResponse, error) {
	return &pb.DeleteChannelsResponse{}, nil
}

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
