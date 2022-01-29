package cache

import "testing"

func TestGetAccount(t *testing.T) {
	TestInitRedis(t)

	GetAccount("hvturinggas")
}
