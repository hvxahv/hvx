package cache

import "testing"

func TestDELAcctMail(t *testing.T) {
	TestInitRedis(t)
	err := DELAcctMail("xxs@disism.com")
	if err != nil {
		return
	}
}
