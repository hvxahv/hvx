package public

import (
	"fmt"
	acct "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	pb "github.com/hvxahv/hvx/api/grpc/proto/public/v1alpha1"
	"github.com/hvxahv/hvx/pkg/activitypub"
	"github.com/hvxahv/hvx/pkg/v"
	"github.com/hvxahv/hvx/pkg/v/cli"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GetPublicInstance(ctx context.Context, g *emptypb.Empty) (*pb.GetPublicInstanceResponse, error) {
	return &pb.GetPublicInstanceResponse{
		Code:       "200",
		Version:    viper.GetString("version"),
		Build:      "2022-01-01",
		Maintainer: viper.GetString("author"),
		Repo:       viper.GetString("name"),
		Host:       viper.GetString("domain"),
	}, nil
}

func (s *server) GetPublicAccountCount(ctx context.Context, g *emptypb.Empty) (*pb.GetPublicAccountCountResponse, error) {
	conn, err := grpc.DialContext(ctx, v.GetGRPCServiceAddress("account"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := cli.NewHvxClient(conn)
	reply, err := client.GetAccountCount(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return &pb.GetPublicAccountCountResponse{
		Code:         "200",
		AccountCount: reply.AccountCount,
	}, nil
}

func (s *server) GetWebfinger(ctx context.Context, in *pb.GetWebfingerRequest) (*pb.GetWebfingerResponse, error) {
	name := activitypub.GetActorName(in.Resource)
	var (
		address = fmt.Sprintf("https://%s/u/%s", viper.GetString("domain"), name)
	)
	conn, err := grpc.DialContext(ctx, v.GetGRPCServiceAddress("account"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := cli.NewHvxClient(conn)

	exist, err := client.IsExist(ctx, &acct.IsExistRequest{
		Username: name,
	})
	if err != nil {
		return nil, err
	}
	if !exist.IsExist {
		return &pb.GetWebfingerResponse{
			Subject: in.Resource,
			Aliases: []string{address},
			Links: []*pb.GetWebfingerResponseLink{
				{
					Rel:  "self",
					Type: "application/activity+json",
					Href: address,
				},
			},
		}, nil
	}
	return &pb.GetWebfingerResponse{
		Subject: "",
		Aliases: nil,
		Links:   nil,
	}, nil
}

func (s *server) GetActor(ctx context.Context, in *pb.GetActorRequest) (*pb.GetActorResponse, error) {
	conn, err := grpc.DialContext(ctx, v.GetGRPCServiceAddress("account"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := cli.NewHvxClient(conn)
	a, err := client.GetActorByAccountUsername(ctx, &acct.GetActorByAccountUsernameRequest{
		Username: in.Actor,
	})
	if err != nil {
		return nil, err
	}

	var (
		address = viper.GetString("domain")

		id  = fmt.Sprintf("https://%s/u/%s", address, a.PreferredUsername)
		kid = fmt.Sprintf("%s#main-key", id)
		box = fmt.Sprintf("https://%s/u/%s/", address, a.PreferredUsername)
	)

	return &pb.GetActorResponse{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1alpha1",
		},
		Id:                id,
		Type:              a.ActorType,
		Following:         "",
		Followers:         "",
		Inbox:             fmt.Sprintf("%s/inbox", box),
		Outbox:            fmt.Sprintf("%s/outbox", box),
		PreferredUsername: a.PreferredUsername,
		Name:              a.Name,
		Summary:           a.Summary,
		Url:               a.Address,
		PublicKey: &pb.GetActorResponse_PublicKey{
			Id:           kid,
			Owner:        a.Address,
			PublicKeyPem: a.PublicKey,
		},
	}, nil
}

func (s *server) CreateAccounts(ctx context.Context, in *pb.CreateAccountsRequest) (*pb.CreateAccountsResponse, error) {
	conn, err := grpc.DialContext(ctx, v.GetGRPCServiceAddress("account"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := cli.NewHvxClient(conn)
	reply, err := client.CreateAccount(ctx, &acct.CreateAccountRequest{
		Username:  in.Username,
		Mail:      in.Mail,
		Password:  in.Password,
		PublicKey: in.PublicKey,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountsResponse{
		Code:     reply.Code,
		Response: reply.Reply,
	}, nil
}

func (s *server) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	conn, err := grpc.DialContext(ctx, v.GetGRPCServiceAddress("account"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := cli.NewHvxClient(conn)
	reply, err := client.Verify(ctx, &acct.VerifyRequest{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticateResponse{
		Code:     reply.Code,
		Token:    reply.Token,
		DeviceId: reply.Id,
	}, nil
}
