package cache

import "testing"

func TestRdb_Dial(t *testing.T) {
	if err := NewRdb().Dial(0); err != nil {
		t.Error(err)
		return
	}
}
