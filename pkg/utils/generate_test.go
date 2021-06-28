package utils

import (
	"fmt"
	"testing"
)

func TestGenPassword(t *testing.T) {
	r := GenPassword("hvxahv")
	t.Log(r)
}

func TestGenRSA(t *testing.T) {
	privateKey, publicKey, err := GenRSA()
	if err != nil {
		return
	}
	fmt.Println(privateKey, publicKey)
}


func TestGenToken(t *testing.T) {
	token, err := GenToken("foo", "bar")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

