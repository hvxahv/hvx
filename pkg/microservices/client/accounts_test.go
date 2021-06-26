package client

import "testing"

func TestAccounts(t *testing.T) {
	accounts, err := Accounts()
	if err != nil {
		return 
	}
	t.Log(accounts)
}
