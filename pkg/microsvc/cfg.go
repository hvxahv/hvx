package microsvc

import (
	"fmt"
	"github.com/spf13/viper"
)

type Option func(o *options)

type options struct {
	id      string
	name    string
	version string
}

type Cfg struct {
	opts *options
}

func WithServiceID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func WithServiceName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}
func WithServiceVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

func New(opts ...Option) *Cfg {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return &Cfg{
		opts: o,
	}
}

type address struct {
	name string
}

func NewAddress(svcName string) *address {
	return &address{
		name: svcName,
	}
}

func GetGRPCServiceAddress(name string) string {
	return fmt.Sprintf("%s:%s", viper.GetString(fmt.Sprintf("microservices.%s.address", name)),
		viper.GetString(fmt.Sprintf("microservices.%s.gp", name)))
}

func GetRestServiceAddress(name string) string {
	return fmt.Sprintf("http://%s:%s", "localhost",
		viper.GetString(fmt.Sprintf("microservices.%s.gwp", name)))
}