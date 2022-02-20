package microservices

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetAccountHost() string {
	return viper.GetString("microservices.account.host")
}

func GetAccountPort() string {
	return viper.GetString("microservices.account.port")
}

func GetAccountAddress() string {
	return fmt.Sprintf("%s:%s", GetAccountHost(), GetAccountPort())
}
