package internal

import (
	"fmt"
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

// Broadcast ...

func (s *server) CreateBroadcast(ctx context.Context, in *pb.CreateBroadcastRequest) (*pb.CreateBroadcastResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	channelId, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}
	articleId, err := strconv.Atoi(in.ArticleId)
	if err != nil {
		return nil, err
	}

	// TODO - SYNC TO IPFS AND RETURN CID.
	// Return CID for simulated IPFS simulation
	cid := uuid.New().String()

	if err := NewBroadcasts(uint(channelId), parse.ActorId, uint(articleId), cid).Create(); err != nil {
		return nil, err
	}
	return &pb.CreateBroadcastResponse{
		Code:  "200",
		Reply: "ok",
	}, nil
}

func (s *server) GetBroadcasts(ctx context.Context, in *pb.GetBroadcastsRequest) (*pb.GetBroadcastsResponse, error) {
	channelId, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}
	broadcasts, err := NewBroadcastsChannelId(uint(channelId)).GetBroadcasts()
	if err != nil {
		return nil, err
	}

	fmt.Println(broadcasts)
	return &pb.GetBroadcastsResponse{
		Code:       "200",
		Broadcasts: nil,
	}, nil

}

func (s *server) DeleteBroadcast(ctx context.Context, in *pb.DeleteBroadcastRequest) (*pb.DeleteBroadcastResponse, error) {
	return &pb.DeleteBroadcastResponse{}, nil
}
