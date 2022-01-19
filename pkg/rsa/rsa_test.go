package rsa

import (
	"fmt"
	"testing"
)

func TestGenRSA(t *testing.T) {
	privateKey, publicKey, err := GenRSA()
	if err != nil {
		return
	}
	fmt.Println(privateKey, publicKey)
}
