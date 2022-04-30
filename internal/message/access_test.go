package message

import (
	"context"
	"fmt"
	"github.com/hvxahv/hvx/api/message/v1alpha1"
	"testing"
)

func TestMessage_MessageAccessRegister(t *testing.T) {
	d := &v1alpha1.MessageAccessRegisterRequest{
		AccountId: "733124680636596225",
		Username:  "hvturingga",
		Password:  "hvxahv123",
	}
	s := message{}
	register, err := s.MessageAccessRegister(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(register)
}

func TestMessage_MessageAccessLogin(t *testing.T) {
	d := &v1alpha1.MessageAccessLoginRequest{
		AccountId: "733124680636596225",
		Username:  "hvturingga",
		Password:  "hvxahv123",
	}
	s := message{}
	login, err := s.MessageAccessLogin(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(login)
}

func TestMessage_MessageAccessDelete(t *testing.T) {
	d := &v1alpha1.MessageAccessDeleteRequest{
		AccountId: "733124680636596225",
	}
	s := message{}
	del, err := s.MessageAccessDelete(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(del)
}
