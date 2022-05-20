package internal

import (
	"fmt"
	"testing"
)

func TestAccount_Verify(t *testing.T) {
	a, err := NewVerify("hvturingga").Verify("hvxahv123")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}
