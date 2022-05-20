package consul

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

type Registry interface {
	Register() error
	Deregister(id string) error
}

func (c *cfg) Register() error {
	cli, err := api.NewClient(&c.consulConfig)
	if err != nil {
		return err
	}

	check := &api.AgentServiceCheck{
		CheckID:                        c.id,
		Name:                           c.name,
		HTTP:                           c.addr,
		Notes:                          fmt.Sprintf("CHECK: %s", c.addr),
		Interval:                       c.interval,
		Timeout:                        c.timeout,
		DeregisterCriticalServiceAfter: c.dcs,
	}

	p, _ := strconv.Atoi(c.port)
	r := &api.AgentServiceRegistration{
		ID:      c.id,
		Name:    c.name,
		Tags:    c.tags,
		Port:    p,
		Address: c.addr,
		Check:   check,
	}

	if err := cli.Agent().ServiceRegister(r); err != nil {
		return errors.Errorf("consul registration failed: %v", err)
	}
	return nil
}

func (c *cfg) Deregister(id string) error {
	cli, err := api.NewClient(&c.consulConfig)
	if err != nil {
		return err
	}

	if err := cli.Agent().ServiceDeregister(id); err != nil {
		return err
	}
	return nil
}
