package account

import (
	"context"
	"fmt"
	"github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"testing"
)

func TestAccount_DHRequestEncryption(t *testing.T) {
	d := &v1alpha1.NewDHRequestEncryption{
		Username:    "hvturingga",
		DeviceId:    "733125490113478657",
		To:          "733125348645240833",
		DhPublicKey: "public_key",
		DhIv:        "iv",
	}

	s := &account{}
	encryption, err := s.DHRequestEncryption(context.Background(), d)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	fmt.Println(encryption)
}

func TestAccount_DHGetPublic(t *testing.T) {
	d := &v1alpha1.NewDHGetPublic{
		DeviceId: "733125348645240833",
	}
	s := account{}
	public, err := s.DHGetPublic(context.Background(), d)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	fmt.Println(public)
}

func TestAccount_DHSendEncryption(t *testing.T) {
	d := &v1alpha1.NewDHSendEncryption{
		DeviceId:    "733125490113478657",
		DhPublicKey: "dh_public_key",
		PrivateKey:  "private_key_aes",
	}
	s := account{}
	encryption, err := s.DHSendEncryption(context.Background(), d)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	fmt.Println(encryption)
}

func TestAccount_DHWaitEncryption(t *testing.T) {
	d := &v1alpha1.NewDHWaitEncryption{
		DeviceId: "733125490113478657",
	}
	s := account{}
	encryption, err := s.DHWaitEncryption(context.Background(), d)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	fmt.Println(encryption)
}
