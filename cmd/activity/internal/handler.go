/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

func (s *server) Inbox(ctx context.Context, in *pb.InboxRequest) (*pb.InboxResponse, error) {
	activity, err := NewActivity(in.Name, in.Data)
	if err != nil {
		return nil, err
	}
	if err := activity.Activity(); err != nil {
		return nil, err
	}
	return &pb.InboxResponse{
		Code:     "200",
		Response: "success",
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
	inbox, err := NewInboxesIdAndActorId(uint(inboxId), parse.ActorId).GetInbox()
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
	if err := NewInboxesIdAndActorId(uint(inboxId), parse.ActorId).DeleteInbox(); err != nil {
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
	inboxes, err := NewInboxesReceiverId(parse.ActorId).GetInboxes()
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
