package main

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/app/accounts/app"
	"log"
)

type server struct {
	pb.AccountsServer

}

// AccountData 将接收到的用户数据存储到数据库
func (s *server) NewAccount(ctx context.Context, in *pb.AccountData) (*pb.NewAccountReply, error) {
	r := app.NewAccount(in)
	log.Println(in)
	i := int32(r)

	return &pb.NewAccountReply{Reply: i}, nil
}

func (s *server) GetAccount(ctx context.Context, in *pb.AccountName) (*pb.AccountData, error) {
	r := app.GetAccountData(in.Username)

	ad := &pb.AccountData{
		Username: r.Username + "这是给你返回的用户名",
	}
	return ad, nil
}