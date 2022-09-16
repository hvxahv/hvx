package internal

import (
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

	// TODO - SYNC TO IPFS AND RETURN CID.
	// Return CID for simulated IPFS simulation
	cid := uuid.New().String()

	if err := NewBroadcasts(uint(in.GetChannelId()), parse.ActorId, uint(in.GetArticleId()), cid).Create(); err != nil {
		return nil, err
	}
	return &pb.CreateBroadcastResponse{
		Code:   "200",
		Status: "ok",
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

	var b []*pb.BroadcastData
	for _, broadcast := range broadcasts {
		b = append(b, &pb.BroadcastData{
			Id:        int64(broadcast.ID),
			ChannelId: int64(broadcast.ChannelId),
			AdminId:   int64(broadcast.AdminId),
			Cid:       broadcast.CID,
		})
	}
	return &pb.GetBroadcastsResponse{
		Code:       "200",
		Broadcasts: b,
	}, nil

}

func (s *server) DeleteBroadcast(ctx context.Context, in *pb.DeleteBroadcastRequest) (*pb.DeleteBroadcastResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewBroadcastsDelete(uint(in.GetBroadcastId()), uint(in.GetChannelId()), parse.ActorId).Delete(); err != nil {
		return nil, err
	}
	return &pb.DeleteBroadcastResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}
