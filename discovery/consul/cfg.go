package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

const (
	address  = "127.0.0.1:8500"
	interval = "10s"
	timeout  = "10s"
	dcs      = "60s"
)

type opts struct {
	consulConfig api.Config // consul default config
	interval     string
	timeout      string
	dcs          string // deregister critical service after.
	q            *api.QueryOptions
}
type cfg struct {
	*opts
	id   string
	name string
	host string
	port string
	addr string
	tags []string
}

type Option func(o *opts)

func SetAddress(addr string) Option {
	return func(o *opts) {
		o.consulConfig.Address = addr
	}
}

func SetWithToken(token string) Option {
	return func(o *opts) {
		o.consulConfig.Token = token
	}
}

func SetWithCheckInterval(interval string) Option {
	return func(o *opts) {
		o.interval = interval
	}
}

func SetWithCheckTimeout(timeout string) Option {
	return func(o *opts) {
		o.timeout = timeout
	}
}

func SetWithDCS(dcs string) Option {
	return func(o *opts) {
		o.dcs = dcs
	}
}

func New(name, host, port string, tags []string, o ...Option) (*cfg, error) {
	df := api.DefaultConfig()
	c := &cfg{
		opts: &opts{
			consulConfig: *df,
		},
		id:   name,
		name: name,
		host: host,
		port: port,
		addr: fmt.Sprintf("http://%s:%s/health", host, port),
		tags: tags,
	}

	for _, i := range o {
		i(c.opts)
	}

	if c.consulConfig.Address == "" {
		c.consulConfig.Address = address
	}
	if c.interval == "" {
		c.interval = interval
	}
	if c.timeout == "" {
		c.timeout = timeout
	}
	if c.dcs == "" {
		c.dcs = dcs
	}
	return c, nil
}

func NewServicesId(id, address string) *cfg {
	ad := api.DefaultConfig()
	ad.Address = address
	return &cfg{
		opts: &opts{
			consulConfig: *ad,
		},
		id: id,
	}
}

func (c *cfg) SetQueryOptions(opts ...Option) *api.QueryOptions {
	o := api.QueryOptions{
		Namespace:         c.q.Namespace,
		Partition:         c.q.Partition,
		Datacenter:        c.q.Datacenter,
		AllowStale:        c.q.AllowStale,
		RequireConsistent: c.q.RequireConsistent,
		UseCache:          c.q.UseCache,
		MaxAge:            c.q.MaxAge,
		StaleIfError:      c.q.StaleIfError,
		WaitIndex:         c.q.WaitIndex,
		WaitHash:          c.q.WaitHash,
		WaitTime:          c.q.WaitTime,
		Token:             c.q.Token,
		Near:              c.q.Near,
		NodeMeta:          c.q.NodeMeta,
		RelayFactor:       c.q.RelayFactor,
		LocalOnly:         c.q.LocalOnly,
		Connect:           c.q.Connect,
		Filter:            c.q.Filter,
	}
	for _, i := range opts {
		i(c.opts)
	}
	return &o
}
