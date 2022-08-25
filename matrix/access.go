package matrix

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/hvxahv/hvx/auth"
)

// https://spec.matrix.org/v1.3/client-server-api/#post_matrixclientv3register
// You need to login and verify the password in the hvx server first,
// and then you will use the account name password to register the matrix account after the verification is successful.
// https://matrix.org/docs/api/#post-/_matrix/client/v3/register
// curl -X POST "https://matrix-client.matrix.org/_matrix/client/v3/register?kind=user" \
// -H "Accept: application/json" \
// -H "Content-Type: application/json" \
// -d '{"auth":{"example_credential":"verypoorsharedsecret","session":"xxxxx","type":"example.type.foo"},"device_id":"GHTYAJCE","inhibit_login":false,"initial_device_display_name":"Jungle Phone","password":"ilovebananas","refresh_token":false,"username":"cheeky_monkey"}' \

type RegisterReq struct {
	Auth struct {
		ExampleCredential string `json:"example_credential"`
		Session           string `json:"session"`
		Type              string `json:"type"`
	} `json:"auth"`
	DeviceId                 string `json:"device_id"`
	InhibitLogin             bool   `json:"inhibit_login"`
	InitialDeviceDisplayName string `json:"initial_device_display_name"`
	Password                 string `json:"password"`
	RefreshToken             bool   `json:"refresh_token"`
	Username                 string `json:"username"`
}

func NewRegisterReq(deviceId, username, password string) *RegisterReq {
	return &RegisterReq{
		Auth: struct {
			ExampleCredential string `json:"example_credential"`
			Session           string `json:"session"`
			Type              string `json:"type"`
		}{
			ExampleCredential: auth.GetTokenSecret(),
			Session:           "",
			Type:              "m.login.dummy",
		},
		DeviceId:                 "",
		InhibitLogin:             false,
		InitialDeviceDisplayName: deviceId,
		Password:                 password,
		RefreshToken:             false,
		Username:                 username,
	}
}

type Res struct {
	Code int
	Body []byte
}

func NewRes(code int, body []byte) *Res {
	return &Res{Code: code, Body: body}
}

func (a *RegisterReq) Register() (*Res, error) {
	// Create a Resty Client
	client := resty.New()

	data, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	address := fmt.Sprintf("https://%s/_matrix/client/v3/register?kind=user", GetMatrixServiceAddress())
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(data).
		Post(address)
	if err != nil {
		return nil, err
	}
	return NewRes(res.StatusCode(), res.Body()), nil
}

// https://spec.matrix.org/v1.3/client-server-api/#post_matrixclientv3accountdeactivate
// Delete this matrix account.
// https://github.com/matrix-org/synapse/issues/1707
//curl -X POST "https://matrix-client.matrix.org/_matrix/client/v3/account/deactivate" \
//-H "Accept: application/json" \
//-H "Content-Type: application/json" \
//-d '{"auth":{"example_credential":"verypoorsharedsecret","session":"xxxxx","type":"example.type.foo"},"id_server":"example.org"}' \

type DeactivateReq struct {
	Auth struct {
		ExampleCredential string `json:"example_credential"`
		Session           string `json:"session"`
		Type              string `json:"type"`
	} `json:"auth"`
	IdServer string `json:"id_server"`
}

func NewDeactivateReq(idServer string) *DeactivateReq {
	return &DeactivateReq{
		Auth: struct {
			ExampleCredential string `json:"example_credential"`
			Session           string `json:"session"`
			Type              string `json:"type"`
		}{
			ExampleCredential: auth.GetTokenSecret(),
			Session:           "",
			Type:              "",
		},
		IdServer: idServer,
	}
}

type DeactivateRes struct {
	IdServerUnbindResult string `json:"id_server_unbind_result"`
}

func (d *DeactivateReq) DeactivateReq() (*Res, error) {
	client := resty.New()

	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s/_matrix/client/v3/account/deactivate", GetMatrixServiceAddress())
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(data).
		Post(address)
	if err != nil {
		return nil, err
	}
	return NewRes(res.StatusCode(), res.Body()), nil
}

// https://spec.matrix.org/v1.3/client-server-api/#post_matrixclientv3accountpassword
// When you change the hvx password, you need to change the matrix account password at the same time.
// https://matrix.org/docs/api/#post-/_matrix/client/v3/account/password
// post /_matrix/client/v3/account/password

type EditPasswordReq struct {
	Auth struct {
		ExampleCredential string `json:"example_credential"`
		Session           string `json:"session"`
		Type              string `json:"type"`
	} `json:"auth"`
	LogoutDevices bool   `json:"logout_devices"`
	NewPassword   string `json:"new_password"`
}

func NewEditPasswordReq(newPassword string, logoutDevices bool) *EditPasswordReq {
	return &EditPasswordReq{
		Auth: struct {
			ExampleCredential string `json:"example_credential"`
			Session           string `json:"session"`
			Type              string `json:"type"`
		}{
			ExampleCredential: auth.GetTokenSecret(),
			Session:           "",
			Type:              "m.login.dummy",
		},
		LogoutDevices: logoutDevices,
		NewPassword:   newPassword,
	}
}

func (e *EditPasswordReq) EditPassword() (*Res, error) {
	client := resty.New()

	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s/_matrix/client/v3/account/password", GetMatrixServiceAddress())
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(data).
		Post(address)
	if err != nil {
		return nil, err
	}
	return NewRes(res.StatusCode(), res.Body()), nil
}
