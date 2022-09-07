package rsa

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	rsa := NewRsa(2048)
	key, err := rsa.Generate()
	if err != nil {
		t.Errorf("generates an RSA keypair error: %v", err)
	}

	if len(key.Private) == 0 {
		t.Errorf("private key is empty")
	}

	if len(key.PublicKey) == 0 {
		t.Errorf("public key is empty")
	}

	//t.Logf("private key: %s", key.Private)
	//t.Logf("public key: %s", key.PublicKey)
	fmt.Println(key.Private)
	fmt.Println(key.PublicKey)

}
