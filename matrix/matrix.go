package matrix

import (
	"encoding/json"
	"github.com/hvxahv/hvx/errors"

	"github.com/go-resty/resty/v2"
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

type MatrixReq struct {
	Address string
	Data    interface{}
}
type Handler interface {
	// Do requesting data from matrix services.
	Do() (*MatrixRes, error)
}

func NewMatrixReq(address string, data interface{}) *MatrixReq {
	return &MatrixReq{Address: address, Data: data}
}

func (m *MatrixReq) Do() (*MatrixRes, error) {
	client := resty.New()

	data, err := json.Marshal(m.Data)
	if err != nil {
		return nil, err
	}
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(data).
		Post(m.Address)
	if err != nil {
		errors.Throw("failed to send a message to the matrix service.", err)
		return nil, err
	}
	return NewMatrixRes(res.StatusCode(), res.Body()), nil
}

type MatrixRes struct {
	Code int
	Body []byte
}

func NewMatrixRes(code int, body []byte) *MatrixRes {
	return &MatrixRes{Code: code, Body: body}
}
