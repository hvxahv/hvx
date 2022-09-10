package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/cmd/activity/internal/outbox"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

//
//func (s *server) CreateOutbox(ctx context.Context, in *pb.CreateOutboxRequest) (*pb.CreateOutboxResponse, error) {
//
//	return &pb.CreateOutboxResponse{}, nil
//}
//
//func (s *server) GetOutbox(ctx context.Context, in *pb.GetOutboxRequest) (*pb.GetOutboxResponse, error) {
//
//	return &pb.GetOutboxResponse{}, nil
//}

func (s *server) GetOutboxes(ctx context.Context, in *emptypb.Empty) (*pb.GetOutboxesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	outboxes, err := outbox.NewOutboxesActorId(parse.ActorId).GetOutboxes()
	if err != nil {
		return nil, err
	}
	var ret []*pb.OutboxData
	for _, obx := range outboxes {
		ret = append(ret, &pb.OutboxData{
			Id:         strconv.Itoa(int(obx.ID)),
			ActorId:    strconv.Itoa(int(obx.ActorId)),
			ActivityId: obx.ActivityId,
			To:         obx.To,
			Cc:         obx.Cc,
			Bcc:        obx.Bcc,
			Bto:        obx.Bto,
			Audience:   obx.Audience,
			Types:      obx.Types,
			Body:       obx.Body,
		})
	}

	return &pb.GetOutboxesResponse{
		Code:     "200",
		Outboxes: ret,
	}, nil
}
