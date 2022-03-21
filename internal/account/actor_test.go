package account

import (
	"context"
	"fmt"
	"testing"

	"github.com/hvxahv/hvxahv/api/account/v1alpha1"
)

func TestAccount_CreateActor(t *testing.T) {
	d := &v1alpha1.CreateActorRequest{
		PreferredUsername: "hvturingga",
		PublicKey:         "idi",
		ActorType:         "Services",
	}
	s := account{}
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
	d := &v1alpha1.GetActorByAccountUsernameRequest{
		Username: "hvxahv2",
	}
	s := account{}
	actor, err := s.GetActorByAccountUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)
}

func TestAccount_GetActorsByPreferredUsername(t *testing.T) {
	d := &v1alpha1.GetActorsByPreferredUsernameRequest{
		PreferredUsername: "hvxahv2",
	}
	s := account{}
	actors, err := s.GetActorsByPreferredUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actors)
}

func TestAccount_GetActorByAddress(t *testing.T) {
	// Local
	d := &v1alpha1.GetActorByAddressRequest{
		Address: "https://6ae9-240e-30d-2200-b600-2db1-88c6-2b4b-ab49.ngrok.io/u/hvturingga",
	}
	s := account{}
	actor, err := s.GetActorByAddress(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)

	// Remote
	d2 := &v1alpha1.GetActorByAddressRequest{
		Address: "https://mastodon.social/users/hvturingga",
	}
	s2 := account{}
	actor2, err := s2.GetActorByAddress(context.Background(), d2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor2)
}

func TestAccount_EditActor(t *testing.T) {
	d := &v1alpha1.EditActorRequest{
		AccountUsername: "hvxahv2",
		Name:            "HVXAHV Test",
		Avatar:          "https://avatars.githubusercontent.com/u/94792300?s=200&v=4",
		Summary:         "",
	}
	s := account{}
	actor, err := s.EditActor(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)
}
