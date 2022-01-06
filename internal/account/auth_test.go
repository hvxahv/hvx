package account

import (
	"fmt"
	"log"
	"testing"
)

func TestAccounts_Login(t *testing.T) {
	a := NewAuthenticate("x@disism.com", "Hvxahv123")
	id, login, err := a.SignIn()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(id)
	fmt.Println(login)
}
