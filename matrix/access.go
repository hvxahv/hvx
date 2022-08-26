package matrix

import (
	"github.com/hvxahv/hvx/auth"
)

// Register https://spec.matrix.org/v1.3/client-server-api/#post_matrixclientv3register
// You need to verify the password in the hvx server first,
// and then you will use the account name password to register the matrix account after the verification is successful.
// https://matrix.org/docs/api/#post-/_matrix/client/v3/register
type Register struct {
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

func NewRegisterReq(deviceId, username, password string) *Register {
	return &Register{
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

// Register Error.
// {
//   "errcode": "M_USER_IN_USE",
//   "error": "Desired user ID is already taken."
// }

type RegisterErrRes struct {
	Errcode string
	Error   string
}

// REGISTER RESPONSE EXAMPLE.
// {
// "access_token": "abc123",
// "device_id": "GHTYAJCE",
// "user_id": "@cheeky_monkey:matrix.org"
// }

type RegisterRes struct {
	AccessToken string `json:"access_token"`
	DeviceId    string `json:"device_id"`
	UserId      string `json:"user_id"`
}

// Deactivate https://spec.matrix.org/v1.3/client-server-api/#post_matrixclientv3accountdeactivate
// Delete this matrix account.
type Deactivate struct {
	Auth struct {
		ExampleCredential string `json:"example_credential"`
		Session           string `json:"session"`
		Type              string `json:"type"`
	} `json:"auth"`
	IdServer string `json:"id_server"`
}

func NewDeactivateReq(idServer string) *Deactivate {
	return &Deactivate{
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

// EditPassword https://spec.matrix.org/v1.3/client-server-api/#post_matrixclientv3accountpassword
// When you change the hvx password, you need to change the matrix account password at the same time.
// https://matrix.org/docs/api/#post-/_matrix/client/v3/account/password
type EditPassword struct {
	Auth struct {
		ExampleCredential string `json:"example_credential"`
		Session           string `json:"session"`
		Type              string `json:"type"`
	} `json:"auth"`
	LogoutDevices bool   `json:"logout_devices"`
	NewPassword   string `json:"new_password"`
}

func NewEditPasswordReq(newPassword string, logoutDevices bool) *EditPassword {
	return &EditPassword{
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
