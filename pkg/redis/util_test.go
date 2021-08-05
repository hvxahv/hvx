package redis

import "testing"

func TestSETAcctMailORUN(t *testing.T) {
	initRedisConfig(t)

	SETAcctMailORUN("x@disism.com", "hvturingga")

}

func TestFINDAcctMailAndUN(t *testing.T) {
	initRedisConfig(t)

	m, u := FINDAcctMailAndUN("x@disism.com", "hvturingg")
	t.Log(m, u)
}