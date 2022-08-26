package matrix

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

// CONFIG EXAMPLE
// matrix:
//   address: "matrix.disism.com"

func GetMatrixServiceAddress() string {
	return viper.GetString("matrix.address")
}

func GetRegisterAddress() string {
	return fmt.Sprintf("https://%s/_matrix/client/v3/register?kind=user", GetMatrixServiceAddress())
}

func GetDeactivateAddress() string {
	return fmt.Sprintf("%s/_matrix/client/v3/account/deactivate", GetMatrixServiceAddress())
}

func GetEditPasswordAddress() string {
	return fmt.Sprintf("%s/_matrix/client/v3/account/password", GetMatrixServiceAddress())
}

type MatrixReq struct {
	Address string
	Data    interface{}
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
