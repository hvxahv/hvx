package internal

import (
	"fmt"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

// Administrative ...

func (s *server) IsAdministrator(ctx context.Context, in *pb.IsAdministratorRequest) (*pb.IsAdministratorResponse, error) {

	return &pb.IsAdministratorResponse{}, nil
}

func (s *server) IsOwner(ctx context.Context, in *pb.IsOwnerRequest) (*pb.IsOwnerResponse, error) {
	return &pb.IsOwnerResponse{}, nil
}

func (s *server) AddAdministrator(ctx context.Context, in *pb.AddAdministratorRequest) (*pb.AddAdministratorResponse, error) {
	return &pb.AddAdministratorResponse{}, nil
}

func (s *server) RemoveAdministrator(ctx context.Context, in *pb.RemoveAdministratorRequest) (*pb.RemoveAdministratorResponse, error) {
	return &pb.RemoveAdministratorResponse{}, nil
}

func (s *server) GetAdministrators(ctx context.Context, in *pb.GetAdministratorsRequest) (*pb.GetAdministratorsResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	channelId, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	accountId, err := strconv.Atoi(parse.AccountId)
	if err != nil {
		return nil, err
	}
	administrators, err := NewAdministratesPermission(uint(channelId), uint(accountId)).GetAdministrators()
	if err != nil {
		return nil, err
	}
	fmt.Println(administrators)
	//var data []*pb.ChannelData
	//for _, d := range administrators {
	//var cd pb.ChannelData
	//client, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")})
	//if err != nil {
	//	return nil, err
	//}
	//defer client.Close()
	//
	//as, err := actor.NewActorClient(client.Conn).Get(ctx, &actor.GetRequest{
	//	ActorId: strconv.Itoa(int(d.)),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//cd.Channels = as.Actor
	//cd.ChannelId = strconv.Itoa(int(d.ID))

	//data = append(data, &cd)
	//}

	return &pb.GetAdministratorsResponse{}, nil
}

func (s *server) ExitAdministrator(ctx context.Context, in *pb.ExitAdministratorRequest) (*pb.ExitAdministratorResponse, error) {
	return &pb.ExitAdministratorResponse{}, nil
}
