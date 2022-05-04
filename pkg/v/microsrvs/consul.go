package microsrvs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strconv"
)

type Register struct {
	ID      string
	Name    string
	Port    string
	Tags    []string
	Address string
}

func NewRegister(name string, port string, tags []string, address string) *Register {
	id := uuid.New().String()
	return &Register{ID: id, Name: name, Port: port, Tags: tags, Address: address}
}

func (r *Register) Register() error {
	cfg := api.DefaultConfig()
	cfg.Address = viper.GetString("consul.address")
	client, err := api.NewClient(cfg)
	if err != nil {
		return errors.Errorf("consul client error : %v", err)
	}

	p, _ := strconv.Atoi(r.Port)
	addr := fmt.Sprintf("http://%s:%d/ping", r.Address, p)
	check := &api.AgentServiceCheck{
		CheckID:                        r.ID,
		Name:                           r.Name,
		HTTP:                           addr,
		Notes:                          fmt.Sprintf("CHECK: %s", addr),
		Interval:                       "30s",
		Timeout:                        "10s",
		DeregisterCriticalServiceAfter: "60s",
	}

	registration := &api.AgentServiceRegistration{
		ID:      r.ID,
		Name:    r.Name,
		Tags:    r.Tags,
		Port:    p,
		Address: r.Address,
		Check:   check,
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		return errors.Errorf("consul registration failed: %v", err)
	}
	return nil
}

func Deregister(id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = viper.GetString("consul.address")

	client, err := api.NewClient(cfg)
	if err != nil {
		return errors.Errorf("consul client error : %v", err)
	}

	if err2 := client.Agent().ServiceDeregister(id); err2 != nil {
		fmt.Println(err2)
		return err2
	}
	return nil
}
