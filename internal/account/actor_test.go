package account

import (
	"context"
	"fmt"
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"testing"
)

func TestAccount_CreateActor(t *testing.T) {
	d := &pb.CreateActorRequest{
		PreferredUsername: "gateway",
		PublicKey:         "idi",
		ActorType:         "Services",
	}
	s := server{}
	r, err := s.CreateActor(context.Background(), d)
	if err != nil {
		t.Error("error:", err)
		return
	}
	if r == nil {
		t.Error("error:", "nil")
		return
	}
	fmt.Println(r)
}

func TestAccount_GetActorByAccountUsername(t *testing.T) {
	d := &pb.GetActorByAccountUsernameRequest{
		Username: "hvxahv2",
	}
	s := server{}
	actor, err := s.GetActorByAccountUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)
}

func TestAccount_GetActorsByPreferredUsername(t *testing.T) {
	d := &pb.GetActorsByPreferredUsernameRequest{
		PreferredUsername: "hvxahv2",
	}
	s := server{}
	actors, err := s.GetActorsByPreferredUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actors)
}

func TestAccount_GetActorByAddress(t *testing.T) {
	// Local
	d := &pb.GetActorByAddressRequest{
		Address: "https://6ae9-240e-30d-2200-b600-2db1-88c6-2b4b-ab49.ngrok.io/u/hvturingga",
	}
	s := server{}
	actor, err := s.GetActorByAddress(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)

	// Remote
	d2 := &pb.GetActorByAddressRequest{
		Address: "https://mastodon.social/users/hvturingga",
	}
	s2 := server{}
	actor2, err := s2.GetActorByAddress(context.Background(), d2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor2)
}

func TestAccount_EditActor(t *testing.T) {
	d := &pb.EditActorRequest{
		AccountUsername: "hvxahv2",
		Name:            "HVXAHV Test",
		Avatar:          "https://avatars.githubusercontent.com/u/94792300?s=200&v=4",
		Summary:         "",
	}
	s := server{}
	actor, err := s.EditActor(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)
}

func TestAccount_DeleteActorByChannelID(t *testing.T) {
	s := server{}
	r, err := s.DeleteActor(context.Background(), &pb.DeleteActorRequest{
		AccountId: "746166817947975681",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r)
}
