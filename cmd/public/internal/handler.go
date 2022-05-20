package internal

import (
	"fmt"
	pb "github.com/hvxahv/hvx/APIs/grpc-go/public/v1alpha1"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) Health(ctx context.Context, g *emptypb.Empty) (*pb.HealthResponese, error) {
	return &pb.HealthResponese{
		Code:   "200",
		Status: "ok",
	}, nil
}

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
	exist, err := NewPublic(ctx).AccountIsExist(name)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if !exist {
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
	a, err := NewPublic(ctx).GetActorByUsername(in.Actor)
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
	res, err := NewPublic(ctx).CreateAccounts(in.Username, in.Password, in.Mail, in.Password)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountsResponse{
		Code:     res.Code,
		Response: res.Reply,
	}, nil
}

func (s *server) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	auth, err := NewPublic(ctx).Auth(in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.AuthenticateResponse{
		Code:     auth.Code,
		Token:    auth.Token,
		DeviceId: auth.Id,
	}, nil
}
