package cache

import "testing"

func TestSETAcctMailORUN(t *testing.T) {
	TestInitRedis(t)

}

func TestSETAcct(t *testing.T) {
	TestInitRedis(t)
	err := SETAcct("hvturingga", []byte("xxsss"), 0)
	if err != nil {
		return
	} else {
		t.Log("ok")
	}
}
