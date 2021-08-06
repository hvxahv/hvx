package cache

import "testing"

func TestFINDAcctMail(t *testing.T) {
	TestInitRedis(t)

	ok := FINDAcctMail("x@disism.com")
	t.Log(ok)
}

func TestExistAcct(t *testing.T) {
	TestInitRedis(t)

	ok := ExistAcct("hvturingga")
	if !ok {
		t.Log("!ok")
	}
}