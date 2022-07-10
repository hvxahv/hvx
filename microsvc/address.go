package microsvc

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewGRPCAddress(servicesName string) string {
	host := viper.GetString(fmt.Sprintf("microsvcs.%s.hostname", servicesName))
	port := viper.GetString(fmt.Sprintf("microsvcs.%s.ports.grpc", servicesName))
	return fmt.Sprintf("%s:%s", host, port)
}
