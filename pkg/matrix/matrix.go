package matrix

import (
	"fmt"
	"github.com/matrix-org/gomatrix"
	"github.com/spf13/viper"
	"log"
)

// RegisterAuth Authentication type at the time of registration.
type RegisterAuth struct {
	ExampleCredential string `json:"example_credential"`
	Session           string `json:"session"`
	Type              string `json:"type"`
}

type Auth struct {
	Username string
	Password string
}

func NewAuth(username string, password string) *Auth {
	return &Auth{Username: username, Password: password}
}

func (a *Auth) Register() (string, error) {
	addr := viper.GetString("matrix.addr")
	cli, err := gomatrix.NewClient(addr, "", "")
	if err != nil {
		return "", err
	}

	rr := gomatrix.ReqRegister{
		Username:                 a.Username,
		BindEmail:                false,
		Password:                 a.Password,
		DeviceID:                 "",
		InitialDeviceDisplayName: "",
		Auth:                     RegisterAuth{
			ExampleCredential: "",
			Session:           "",
			Type:              "m.login.dummy",
		},
	}

	register, r, err := cli.Register(&rr)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println(r)
	return register.UserID, nil
}