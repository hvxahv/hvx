package activitypub

import (
	"fmt"
	"testing"
)

func TestNewWebFinger(t *testing.T) {
	wf := NewWebFinger("hvturingga")
	fmt.Println(wf)
}

func TestGetWebFinger(t *testing.T) {
	GetWebFinger("hvturingga@mas.to")
}