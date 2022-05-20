package cache

import "testing"

func TestRdb_Dial(t *testing.T) {
	NewRdb().Dial(0)
}
