package main

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/app/inbox/services"
	"log"
)

/**
Inbox 服务端的 handler
*/

// NewInbox 将接收到的收件箱的新消息保存并通知用户
func (s *server) NewInbox(ctx context.Context, in *pb.InboxData) (*pb.NewInboxReply, error) {
	r := services.ReceiveInbox(in)

	return &pb.NewInboxReply{Reply: r}, nil
}


// NewInbox 将接收到的收件箱的新消息保存并通知用户
func (s *server) GetInbox(ctx context.Context, in *pb.Name) (*pb.GetInboxReply, error) {
	r := services.GetInboxData(in)

	log.Println("获得的 INBOX 数据, ", r)
	return &pb.GetInboxReply{Inbox: r}, nil
}

