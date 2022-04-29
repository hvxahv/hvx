package v

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

func GetServiceAddresses(name string) string {
	return viper.GetString(fmt.Sprintf("microservices.%s.port", name))
}

func GetServiceEndpoint(name string) string {
	h := viper.GetString(fmt.Sprintf("microservices.%s.host", name))
	return fmt.Sprintf("%s:%s", h, GetServiceAddresses(name))
}

func GetRestEndpoint(name string) string {
	return viper.GetString(fmt.Sprintf("microservices.%s.gw", name))
}
