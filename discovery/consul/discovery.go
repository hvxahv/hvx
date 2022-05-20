package consul

import (
	"github.com/hashicorp/consul/api"
)

type discovery interface {
	FindServices() (map[string]*api.AgentService, error)
	FindServiceById() (*api.AgentService, *api.QueryMeta, error)
}

func NewDiscovery(address string) *cfg {
	ad := api.DefaultConfig()
	ad.Address = address
	return &cfg{
		opts: &opts{
			consulConfig: *ad,
		},
	}
}

func (c *cfg) FindServices() (map[string]*api.AgentService, error) {
	cli, err := api.NewClient(&c.consulConfig)
	if err != nil {
		return nil, err
	}

	svc, err := cli.Agent().Services()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (c *cfg) FindServiceById() (*api.AgentService, *api.QueryMeta, error) {
	cli, err := api.NewClient(&c.consulConfig)
	if err != nil {
		return nil, nil, err
	}

	svc, qm, err := cli.Agent().Service(c.id, nil)
	if err != nil {
		return nil, nil, err
	}
	return svc, qm, err
}
