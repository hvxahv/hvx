package internal

import (
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/public"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func (s *server) GetInstance(ctx context.Context, g *emptypb.Empty) (*pb.GetInstanceResponse, error) {
	return &pb.GetInstanceResponse{
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
	exist, err := NewPublic(ctx).IsExist(name)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// If the user exists, Webfinger data will be returned.
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
		Inbox:             a.Inbox,
		Outbox:            fmt.Sprintf("%soutbox", box),
		PreferredUsername: a.PreferredUsername,
		Name:              a.Name,
		Summary:           a.Summary,
		Url:               a.Address,
		PublicKey: &pb.GetActorResponse_PublicKey{
			Id:           kid,
			Owner:        a.Address,
			PublicKeyPem: a.PublicKey,
		},
		Icon: &pb.GetActorResponse_Icon{
			Type:      "Image",
			MediaType: "image/jpeg",
			Url:       a.Avatar,
		},
	}, nil
}

func (s *server) CreateAccounts(ctx context.Context, in *pb.CreateAccountsRequest) (*pb.CreateAccountsResponse, error) {
	res, err := NewPublic(ctx).CreateAccount(in.Username, in.Mail, in.Password, in.PublicKey)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountsResponse{
		Code:     res.Code,
		Response: res.Reply,
	}, nil
}
