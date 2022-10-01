package matrix

import (
	"maunium.net/go/mautrix"
)

// Currently, due to the numerous server implementations of matrix.org, I am experiencing
// unsuccessful registration after deploying the https://gitlab.com/famedly/conduit server.
// Since the matrix package of hvx has not been implemented in detail and tested extensively,
// I would like to use a more mature and stable library to implement matrix. org protocol's account features.
// Then I used js sdk to register without any problems, so hvx will decide to move all matrix.org
// account operations to client-side implementation, and the server will accept the data after
// successful client-side operations for saving. So this package will be deprecated in the future.

// https://matrix.org/
// An open network for secure, decentralized communication.
// https://matrix.org/docs/api/#overview
// https://spec.matrix.org/v1.3/
// https://spec.matrix.org/v1.2/client-server-api/#post_matrixclientv3register

type Matrix struct {
	Client *mautrix.Client
	Err    error
}

func New(address string) *Matrix {
	client, err := mautrix.NewClient(address, "", "")
	if err != nil {
		return &Matrix{Client: nil, Err: err}
	}
	return &Matrix{Client: client, Err: nil}
}
