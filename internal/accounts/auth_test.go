package accounts

import (
	"fmt"
	"log"
	"testing"
)

func TestAccounts_Login(t *testing.T) {
	TestInitDB(t)

	a := NewAuth("x@disism.com", "Hvxahv123")
	login, err := a.Login()
	if err != nil {
		log.Println(err)
		return 
	}

	fmt.Println(login)
}
