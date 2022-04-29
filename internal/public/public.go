package public

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/public/v1alpha1"
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

func (s *server) GetPublicInstance(ctx context.Context, g *emptypb.Empty) (*pb.GetPublicInstanceResponse, error) {
	return &pb.GetPublicInstanceResponse{
		Code:       "200",
		Version:    viper.GetString("version"),
		Build:      "2022-01-01",
		Maintainer: viper.GetString("author"),
		Repo:       viper.GetString("name"),
		Host:       viper.GetString("localhost"),
	}, nil
}

func (s *server) GetPublicAccountCount(ctx context.Context, g *emptypb.Empty) (*pb.GetPublicAccountCountResponse, error) {
	db := cockroach.GetDB()
	var count int64
	if err := db.Debug().Table("accounts").Count(&count).Error; err != nil {
		return nil, err
	}
	return &pb.GetPublicAccountCountResponse{
		Code:         "200",
		AccountCount: strconv.Itoa(int(count)),
	}, nil
}

func (s *server) GetWebfinger(ctx context.Context, in *pb.GetWebfingerRequest) (*pb.GetWebfingerResponse, error) {
	return &pb.GetWebfingerResponse{
		Context:   nil,
		Id:        "",
		Type:      "",
		PublicKey: nil,
	}, nil
}
