package notify

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	if err := NewPush(720140086112714753, []byte("Life's Not Out To Get You")).Push(); err != nil {
		fmt.Println(err)
		return
	}
}
