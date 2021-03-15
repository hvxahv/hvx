package main

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/app/outbox/services"
	"hvxahv/pkg/outbox"
	"log"
)

// Accept 同意接收到的请求
func (s *server) Accept(ctx context.Context, in *pb.AcceptData) (*pb.ReplyCode, error) {
	d := inbox.NewAccept(in.Actor, in.Name, in.RequestId)
	r := services.AcceptHandler(d)
	log.Println("---- Accept 服务端接收到的消息 --> ", in.Actor, in.Name, in.RequestId)
	return &pb.ReplyCode{Reply: int32(r)}, nil
}
// Follow 请求关注
func (s *server) Follow(ctx context.Context, in *pb.FollowData) (*pb.ReplyCode, error) {
	r := services.FollowHandler(in.Actor, in.Name)
	return &pb.ReplyCode{Reply: int32(r)}, nil
}

