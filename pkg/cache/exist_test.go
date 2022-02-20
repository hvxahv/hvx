package cache

import (
	"fmt"
	"testing"
)

func TestFINDAcctMail(t *testing.T) {
	TestInitRedis(t)

	ok := SISAcctMail("x@disism.com")
	t.Log(ok)
}

func TestExistAcct(t *testing.T) {
	TestInitRedis(t)

	ok := SISAcct("xxs")
	if !ok {
		t.Log("!ok")
	}
	fmt.Println(ok)
}
