package client

import "testing"

func TestAccounts(t *testing.T) {
	_, _, accounts := Accounts()
	if err != nil {
		return 
	}
	t.Log(accounts)
}
