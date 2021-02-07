package main

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/app/outbox/services"
	"hvxahv/pkg/models"
	"log"
)

// Accept 同意接收到的请求
func (s *server) Accept(ctx context.Context, in *pb.AcceptData) (*pb.AcceptReply, error) {
	d := models.NewAccept(in.Actor, in.Name, in.RequestId)
	services.AcceptHandler(d)
	log.Println("---- Accept 服务端接收到的消息 --> ", in.Actor, in.Name, in.RequestId)
	return &pb.AcceptReply{Reply: "ok!"}, nil
}
