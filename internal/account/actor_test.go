package account

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"golang.org/x/net/context"
	"testing"
)

func TestActors_GetActorByAccountUsername(t *testing.T) {
	n := &pb.NewAccountUsername{Username: "hvxahv"}
	s := &account{}
	a, err := s.GetActorByAccountUsername(context.Background(), n)
	if err != nil {
		t.Errorf("Error getting actors: %v", err)
		return
	}
	fmt.Println(a)
}

func TestAccount_GetActorsByPreferredUsername(t *testing.T) {
	n := &pb.NewActorPreferredUsername{PreferredUsername: "hvxahv"}
	s := &account{}
	a, err := s.GetActorsByPreferredUsername(context.Background(), n)
	if err != nil {
		t.Errorf("Error getting actors: %v", err)
		return
	}
	fmt.Println(a.Code, a.Actors)
}

func TestAccount_AddActor(t *testing.T) {
	actor := &pb.ActorData{
		PreferredUsername: "xxs",
		Domain:            "https://halfmemories.com",
		Avatar:            "https://www.halfmemories.com/avatar.png",
		Name:              "HVTURINGGA",
		Summary:           "Aha...",
		Inbox:             "https://halfmemories.com/u/xxs/inbox",
		Address:           "https://halfmemories.com/u/xxs",
		PublicKey:         "public_key",
		ActorType:         "Person",
	}
	s := &account{}
	a, err := s.AddActor(context.Background(), actor)
	if err != nil {
		t.Errorf("Error adding actor: %v", err)
		return
	}
	fmt.Println(a.Code, a.Reply)
}

func TestAccount_UpdateActor(t *testing.T) {
	actor := &pb.NewEditActor{
		AccountUsername: "hvxahv",
		Avatar:          "https://www.halfmemories.com/kobayashi.png",
		Name:            "KOBAYASHI YUI",
		Summary:         "Bio...",
	}
	s := &account{}
	a, err := s.EditActor(context.Background(), actor)
	if err != nil {
		t.Errorf("Error updating actor: %v", err)
		return
	}
	fmt.Println(a.Code, a.Reply)
}
