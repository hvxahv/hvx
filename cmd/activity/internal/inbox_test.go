package internal

import "testing"

func TestInboxes_Create(t *testing.T) {
	if err := NewInboxes(1, 2, "3", "4", []byte("5")).Create(); err != nil {
		t.Error(err)
		return
	}
}
