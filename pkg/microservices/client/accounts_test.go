package client

import (
	"testing"
)

func TestAccounts(t *testing.T) {
	_, err, accounts := Accounts()
	if err != nil {
		return 
	}
	t.Log(accounts)
}
