package chat

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/matrix"
	"github.com/matrix-org/gomatrix"
	"log"
)

type Register struct {
	ExampleCredential string `json:"example_credential"`
	Session           string `json:"session"`
	Type              string `json:"type"`
}

type RegisterAuth struct {
	AccountID uint
	Username  string
	Password  string
}

func (a *RegisterAuth) Register() error {
	cli, err := matrix.NewClient("", "")
	if err != nil {
		return err
	}

	rr := gomatrix.ReqRegister{
		Username:                 a.Username,
		BindEmail:                false,
		Password:                 a.Password,
		DeviceID:                 "",
		InitialDeviceDisplayName: "",
		Auth: Register{
			ExampleCredential: "",
			Session:           "",
			Type:              "m.login.dummy",
		},
	}

	register, r, err := cli.Register(&rr)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(r)
	c := NewMatrixAccess(a.AccountID, register.AccessToken, register.HomeServer, register.UserID, register.DeviceID)
	if err := c.Create(); err != nil {
		return err
	}
	return nil
}

func NewMatrixAuth(accountID uint, username string, password string) *RegisterAuth {
	return &RegisterAuth{AccountID: accountID, Username: username, Password: password}
}

type Authentication interface {
	Register() error
}
