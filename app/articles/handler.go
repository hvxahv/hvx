package main

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/kernel/v1"
	"log"
)

type server struct {
	pb.ArticlesServer
}
// NewStatus 新建一篇状态
func (s *server) NewArticle(ctx context.Context, in *pb.ArticleData) (*pb.NewArticleReply, error) {
	log.Println("得到的新文章数据",in)
	return &pb.NewArticleReply{Reply: "创建文章成功"}, nil
}

// UpdateStatus 修改状态
func (s *server) UpdateArticle(ctx context.Context, in *pb.ArticleData) (*pb.UpdateArticleReply, error) {
	log.Println("拿到的更新文章的 ID", in.Id)
	log.Println("更新文章的数据",in)
	return &pb.UpdateArticleReply{Reply: "更新文章成功"}, nil
}

// DeleteStatus 删除状态（根据 ID）
func (s *server) DeleteArticle(ctx context.Context, in *pb.DeleteArticleByID) (*pb.DeleteArticleReply, error) {
	log.Println("得到的删除 ID",in)
	return &pb.DeleteArticleReply{Reply: "删除状态成功"}, nil
}

