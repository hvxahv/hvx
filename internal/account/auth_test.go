package account

import (
	"context"
	"fmt"
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"testing"
)

func TestAccount_Verify(t *testing.T) {
	d := &pb.VerifyRequest{
		Username: "hvxahv2",
		Password: "hvxahv1234",
		Ua:       "chrome",
	}
	s := server{}
	verify, err := s.Verify(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(verify)
}

func TestAccount_GetPublicKeyByAccountUsername(t *testing.T) {
	d := &pb.GetPublicKeyByAccountUsernameRequest{
		Username: "hvxahv2",
	}
	s := server{}
	publicKey, err := s.GetPublicKeyByAccountUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(publicKey)
}
