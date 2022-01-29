package account

import (
	"fmt"
	"github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"golang.org/x/net/context"
	"testing"
)

func TestAccount_Verify(t *testing.T) {
	d := &v1alpha1.NewAccountVerify{
		Username: "hvxahv",
		Password: "Hvxahv123",
		Ua:       "",
	}
	s := &account{}
	r, err := s.Verify(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r.Code, r.Reply)
}

func TestAccount_GetPublicKeyByAccountUsername(t *testing.T) {
	d := &v1alpha1.NewAccountUsername{
		Username: "hvxahv",
	}
	s := &account{}
	r, err := s.GetPublicKeyByAccountUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r.Code, r.PublicKey)
}
