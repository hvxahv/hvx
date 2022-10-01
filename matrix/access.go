package matrix

import (
	"maunium.net/go/mautrix"
)

// Register https://spec.matrix.org/v1.3/client-server-api/#post_matrixclientv3register
// You need to verify the password in the hvx server first,
// and then you will use the account name password to register the matrix account after the verification is successful.
// https://matrix.org/docs/api/#post-/_matrix/client/v3/register

type Access interface {
	RegisterDummy(username, password string) (*mautrix.RespRegister, error)
}

func (x *Matrix) RegisterDummy(username, password string) (*mautrix.RespRegister, error) {
	res, err := x.Client.RegisterDummy(&mautrix.ReqRegister{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
