package main
//
//import (
//	"golang.org/x/net/context"
//	pb "hvxahv/api/hvxahv/v1"
//	"hvxahv/services/accounts.bac/services"
//	"log"
//)
//
//type server struct {
//	pb.AccountsServer
//
//}
//
//// NewAccount 创建账户 将接收到的用户数据存储到数据库
//func (s *server) NewAccount(ctx context.Context, in *pb.AccountData) (*pb.NewAccountReply, error) {
//	r := services.NewAccount(in)
//	log.Println(in)
//	i := int32(r)
//
//	return &pb.NewAccountReply{Reply: i}, nil
//}
//
//// GetAccount 获取账户资料
//func (s *server) GetAccount(ctx context.Context, in *pb.AccountName) (*pb.AccountData, error) {
//	r := services.GetAccountData(in.Username)
//
//	ad := &pb.AccountData{
//		Username: r.Username,
//	}
//	return ad, nil
//}