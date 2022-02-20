package microservices

import (
	"fmt"
	"github.com/spf13/viper"
)

// GetAccountHost returns the hostname of the account microservice.
func GetAccountHost() string {
	return viper.GetString("microservices.account.host")
}

// GetAccountPort returns the port of the account microservice.
func GetAccountPort() string {
	return viper.GetString("microservices.account.port")
}

// GetAccountAddress returns the URL of the account microservice.
func GetAccountAddress() string {
	return fmt.Sprintf("%s:%s", GetAccountHost(), GetAccountPort())
}

// GetSavedHost returns the host of the saved microservice.
func GetSavedHost() string {
	return viper.GetString("microservices.saved.port")
}

// GetSavedPort returns the port of the saved microservice.
func GetSavedPort() string {
	return viper.GetString("microservices.saved.port")
}

// GetSavedAddress returns the URL of the saved microservice.
func GetSavedAddress() string {
	return fmt.Sprintf("%s:%s", GetSavedHost(), GetSavedPort())
}

func GetMessageHost() string {
	return viper.GetString("microservices.message.host")
}

func GetMessagePort() string {
	return viper.GetString("microservices.message.port")
}

func GetMessageAddress() string {
	return fmt.Sprintf("%s:%s", GetMessageHost(), GetMessagePort())
}
