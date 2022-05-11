package public

import (
	"fmt"
	"time"

	acct "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	pb "github.com/hvxahv/hvx/api/grpc/proto/public/v1alpha1"
	"github.com/hvxahv/hvx/pkg/activitypub"
	"github.com/hvxahv/hvx/pkg/microsvc"
	clientv1 "github.com/hvxahv/hvx/pkg/microsvc/client/v1"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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

func (s *server) GetWebfinger(ctx context.Context, in *pb.GetWebfingerRequest) (*pb.GetWebfingerResponse, error) {
	name := activitypub.GetActorName(in.Resource)
	var (
		address = fmt.Sprintf("https://%s/u/%s", viper.GetString("domain"), name)
	)

	cli, err := clientv1.New(ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialOptionsWithToken(),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	exist, err := cli.IsExist(ctx, &acct.IsExistRequest{
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
	cli, err := clientv1.New(ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialOptionsWithToken(),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	a, err := cli.GetActorByUsername(ctx, &acct.GetActorByUsernameRequest{
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
	cli, err := clientv1.New(ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialOptionsWithToken(),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	reply, err := cli.CreateAccount(ctx, &acct.CreateAccountRequest{
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
	cli, err := clientv1.New(ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialOptionsWithToken(),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	reply, err := cli.Verify(ctx, &acct.VerifyRequest{
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
