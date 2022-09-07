package matrix

import (
	"fmt"
	"github.com/spf13/viper"
)

// CONFIG EXAMPLE
// matrix:
//   address: "matrix.disism.com"

const (
	matrixDefaultAddress = "matrix.disism.com"
)

func GetMatrixAddress() string {
	address := viper.GetString("matrix.address")
	if address != "" {
		return address
	}
	return matrixDefaultAddress
}

// GetRegisterAddress EXAMPLE: https://matrix.org/docs/api/#post-/_matrix/client/v3/register
func GetRegisterAddress() string {
	return fmt.Sprintf("%s/_matrix/client/v3/register?kind=user", GetMatrixAddress())
}

func GetDeactivateAddress() string {
	return fmt.Sprintf("%s/_matrix/client/v3/account/deactivate", GetMatrixAddress())
}

func GetEditPasswordAddress() string {
	return fmt.Sprintf("%s/_matrix/client/v3/account/password", GetMatrixAddress())
}
