package matrix

import "github.com/spf13/viper"

// CONFIG EXAMPLE
// matrix:
//   address: "matrix.disism.com"

func GetMatrixServiceAddress() string {
	return viper.GetString("matrix.address")
}
