package internal

import (
	"fmt"
	"strconv"
	"testing"

	pb "github.com/hvxahv/hvx/APIs/grpc-go/account/v1alpha1"
	"github.com/hvxahv/hvx/conv"
)

func TestAccount_GetActorsByPreferredUsername(t *testing.T) {
	actors, err := NewPreferredUsername("hvx1").GetActorsByPreferredUsername()
	if err != nil {
		t.Error(err)
		return
	}

	var a []*pb.ActorDataResponse
	for _, v := range actors {
		var ad *pb.ActorDataResponse
		ad.Id = conv.UintToString(v.ID)
		ad.PreferredUsername = v.PreferredUsername
		ad.Domain = v.Domain
		ad.Avatar = v.Avatar
		ad.Name = v.Name
		ad.Summary = v.Summary
		ad.Inbox = v.Inbox
		ad.Address = v.Address
		ad.PublicKey = v.PublicKey
		ad.ActorType = v.ActorType
		ad.IsRemote = strconv.FormatBool(v.IsRemote)

		a = append(a, ad)
	}
	fmt.Println(a)
}

func TestAccount_GetActorByAddress(t *testing.T) {
	a, err := NewActorAddress("").GetActorByAddress()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)

	// Remote
	a2, err := NewActorAddress("https://mastodon.social/users/hvturingga").GetActorByAddress()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a2)
}

func TestAccount_EditActor(t *testing.T) {
	a := Actors{}
	a.SetActorAvatar("")
	a.SetActorName("")
	a.SetActorSummary("")

	const accountId uint = 1234567890
	if err := a.EditActor(accountId); err != nil {
		t.Error(err)
		return
	}
}
