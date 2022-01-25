package account

import (
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"golang.org/x/net/context"
)

func (s *server) Create(ctx context.Context, in *pb.CreateData) (*pb.CreateReply, error) {
	if err := NewAccounts(in.Username, in.Mail, in.Password).Create(); err != nil {
		return &pb.CreateReply{Code: "202", Reply: err.Error()}, nil
	}

	_, err := NewActors(in.Username, in.PublicKey, "Person").Create()
	if err != nil {
		return &pb.CreateReply{Code: "202", Reply: err.Error()}, nil
	}

	return &pb.CreateReply{Code: "200", Reply: "ok"}, nil
}
