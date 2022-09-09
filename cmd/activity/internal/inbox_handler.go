package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	inbox2 "github.com/hvxahv/hvx/cmd/activity/internal/inbox"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

func (s *server) Inbox(ctx context.Context, in *pb.InboxRequest) (*pb.InboxResponse, error) {
	activity, err := inbox2.NewActivity(in.Name, in.Data)
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
	inbox, err := inbox2.NewInboxesIdAndActorId(uint(inboxId), parse.ActorId).GetInbox()
	if err != nil {
		return nil, err
	}
	return &pb.GetInboxResponse{
		Code: "200",
		Inbox: &pb.Inboxes{
			Id:           in.GetInboxId(),
			ReceiverId:   strconv.Itoa(int(inbox.ReceiverId)),
			SenderAddr:   inbox.SenderAddr,
			ActivityId:   inbox.ActivityId,
			ActivityType: inbox.ActivityType,
			ActivityBody: inbox.ActivityBody,
		},
	}, nil
}

func (s *server) DeleteInbox(ctx context.Context, in *pb.DeleteInboxRequest) (*pb.DeleteInboxResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	inboxId, err := strconv.Atoi(in.InboxId)
	if err != nil {
		return nil, err
	}
	if err := inbox2.NewInboxesIdAndActorId(uint(inboxId), parse.ActorId).DeleteInbox(); err != nil {
		return nil, err
	}
	return &pb.DeleteInboxResponse{
		Code:  "200",
		Reply: "ok",
	}, nil
}

func (s *server) GetInboxes(ctx context.Context, in *emptypb.Empty) (*pb.GetInboxesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	inboxes, err := inbox2.NewInboxesReceiverId(parse.ActorId).GetInboxes()
	if err != nil {
		return nil, err
	}
	var ret []*pb.Inboxes
	for _, inbox := range inboxes {
		ret = append(ret, &pb.Inboxes{
			Id:           strconv.Itoa(int(inbox.ID)),
			ReceiverId:   strconv.Itoa(int(inbox.ReceiverId)),
			SenderAddr:   inbox.SenderAddr,
			ActivityId:   inbox.ActivityId,
			ActivityType: inbox.ActivityType,
			ActivityBody: inbox.ActivityBody,
		})
	}

	return &pb.GetInboxesResponse{
		Code:    "200",
		Inboxes: ret,
	}, nil
}
