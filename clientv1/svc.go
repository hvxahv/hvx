package clientv1

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

type Svc struct {
	ctx     context.Context
	address string
}

func New(ctx context.Context, serviceName string) *Svc {
	return &Svc{
		ctx:     ctx,
		address: GetServiceAddress(serviceName),
	}
}

// GetServiceAddress receives the service name and returns the gRPC address. config provider is viper.
func GetServiceAddress(servicesName string) (address string) {
	var (
		host = viper.GetString(fmt.Sprintf("microsvcs.%s.hostname", servicesName))
		port = viper.GetString(fmt.Sprintf("microsvcs.%s.ports.grpc", servicesName))
	)
	addr := fmt.Sprintf("%s:%s", host, port)
	return addr
}
