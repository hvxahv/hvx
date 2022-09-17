package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/cmd/activity/internal/inbox"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

func (s *server) Inbox(ctx context.Context, in *pb.InboxRequest) (*pb.InboxResponse, error) {
	activity, err := inbox.NewActivity(in.Name, in.Data)
	if err != nil {
		return nil, err
	}

	if err := activity.Handler(); err != nil {
		return nil, err
	}
	return &pb.InboxResponse{
		Code:   "200",
		Status: "success",
	}, nil
}

func (s *server) GetInbox(ctx context.Context, in *pb.GetInboxRequest) (*pb.GetInboxResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	inboxId, err := strconv.Atoi(in.InboxId)
	if err != nil {
		return nil, err
	}

	ibx, err := inbox.NewInboxesIdAndActorId(uint(inboxId), parse.ActorId).GetInbox()
	if err != nil {
		return nil, err
	}
	return &pb.GetInboxResponse{
		Code: "200",
		Inbox: &pb.Inboxes{
			Id:         in.GetInboxId(),
			ActorId:    strconv.Itoa(int(ibx.ActorId)),
			From:       ibx.From,
			ActivityId: ibx.ActivityId,
			Type:       ibx.Types,
			Body:       ibx.Body,
		},
	}, nil
}

func (s *server) DeleteInbox(ctx context.Context, in *pb.DeleteInboxRequest) (*pb.DeleteInboxResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	if err := inbox.NewInboxesIdAndActorId(uint(in.GetInboxId()), parse.ActorId).DeleteInbox(); err != nil {
		return nil, err
	}
	return &pb.DeleteInboxResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) GetInboxes(ctx context.Context, in *emptypb.Empty) (*pb.GetInboxesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	inboxes, err := inbox.NewInboxesReceiverId(parse.ActorId).GetInboxes()
	if err != nil {
		return nil, err
	}
	var ret []*pb.Inboxes
	for _, ibx := range inboxes {
		ret = append(ret, &pb.Inboxes{
			Id:         strconv.Itoa(int(ibx.ID)),
			ActorId:    strconv.Itoa(int(ibx.ActorId)),
			From:       ibx.From,
			ActivityId: ibx.ActivityId,
			Type:       ibx.Types,
			Body:       ibx.Body,
			Viewed:     ibx.Viewed,
		})
	}

	return &pb.GetInboxesResponse{
		Code:    "200",
		Inboxes: ret,
	}, nil
}

func (s *server) ViewedInbox(ctx context.Context, in *pb.ViewedInboxRequest) (*pb.ViewedInboxResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := inbox.NewSetViewed(parse.ActorId, uint(in.GetInboxId())).SetViewed(); err != nil {
		return nil, err
	}
	return &pb.ViewedInboxResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}
