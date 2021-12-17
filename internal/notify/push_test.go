package notify

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	if err := NewPush(2131312412213312312, []byte("Life's Not Out To Get You")).Push(); err != nil {
		fmt.Println(err)
		return
	}
}
