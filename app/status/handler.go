package main

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/kernel/v1"
	"log"
)

type server struct {
	pb.StatusServer
}
// NewStatus 新建一篇状态
func (s *server) NewStatus(ctx context.Context, in *pb.StatusData) (*pb.NewStatusReply, error) {
	log.Println("得到的新状态数据",in)
	return &pb.NewStatusReply{Reply: "新建状态成功"}, nil
}

// UpdateStatus 修改状态
func (s *server) UpdateStatus(ctx context.Context, in *pb.StatusData) (*pb.UpdateStatusReply, error) {
	log.Println("拿到的更新状态的 ID", in.Id)
	log.Println("更新状态的数据",in)
	return &pb.UpdateStatusReply{Reply: "更新状态成功"}, nil
}

// DeleteStatus 删除状态（根据 ID）
func (s *server) DeleteStatus(ctx context.Context, in *pb.DeleteStatusByID) (*pb.DeleteStatusReply, error) {
	log.Println("得到的删除 ID",in)
	return &pb.DeleteStatusReply{Reply: "删除状态成功"}, nil
}
