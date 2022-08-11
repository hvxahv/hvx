package microsvc

import (
	"fmt"

	"github.com/spf13/viper"
)

type Address struct {
	Host    string
	Port    string
	Address string
}

type Addresses interface {
	Get() (address string)
}

func NewGRPCAddress(servicesName string) *Address {
	host := viper.GetString(fmt.Sprintf("microsvcs.%s.hostname", servicesName))
	port := viper.GetString(fmt.Sprintf("microsvcs.%s.ports.grpc", servicesName))
	return &Address{
		Host: host,
		Port: port,
	}
}

func (a *Address) Get() (address string) {
	addr := fmt.Sprintf("%s:%s", a.Host, a.Port)
	return addr
}
