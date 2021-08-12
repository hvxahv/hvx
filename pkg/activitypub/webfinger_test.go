package activitypub

import (
	"fmt"
	"testing"
)

func TestNewWebFinger(t *testing.T) {
	wf := NewWebFinger("hvturingga")
	fmt.Println(wf)
}
