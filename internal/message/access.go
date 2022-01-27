package message

import (
	"log"

	"github.com/hvxahv/hvxahv/pkg/matrix"
	"github.com/matrix-org/gomatrix"
)

type Register struct {
	ExampleCredential string `json:"example_credential"`
	Session           string `json:"session"`
	Type              string `json:"type"`
}

type MatrixAccess struct {
	AccountID uint
	Username  string
	Password  string
}

func (a *MatrixAccess) Login() error {
	cli, err := matrix.NewClient(a.Username, a.Password)
	if err != nil {
		return err
	}

	r := &gomatrix.ReqLogin{
		Password: a.Password,
		User:     a.Username,
	}

	login, err := cli.Login(r)
	if err != nil {
		return err
	}

	if err := NewAccessUpdateTokenByAcctID(a.AccountID, login.AccessToken).UpdateToken(); err != nil {
		return err
	}
	return nil
}

func (a *MatrixAccess) Register() error {
	cli, err := matrix.NewClient("", "")
	if err != nil {
		return err
	}

	reg := &gomatrix.ReqRegister{
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

	resp, _, err := cli.Register(reg)
	if err != nil {
		log.Println(err)
		return err
	}

	c := NewMatrixAccesses(a.AccountID, resp.AccessToken, resp.HomeServer, resp.UserID, resp.DeviceID)
	if err := c.Create(); err != nil {
		return err
	}
	return nil
}

//func NewMatrixAccessAuth(username string, password string) *MatrixAccess {
//	a, err := account.NewAccountsUsername(username).GetAccountByUsername()
//	if err != nil {
//		return nil
//	}
//	return &MatrixAccess{AccountID: a.ID, Username: username, Password: password}
//}

type Authentication interface {
	// Register When registering an account, you need to call this method to register an account of the matrix protocol.
	Register() error
	// Login If the login expires, you need to log in again with a password, and update the token in the matrix_accesses table after the login is completed.
	Login() error
}
