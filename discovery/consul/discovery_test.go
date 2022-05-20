package consul

import (
	"fmt"
	"testing"
)

func TestCfg_FindServices(t *testing.T) {
	n := NewDiscovery("10.143.24.84:8500")
	services, err := n.FindServices()
	if err != nil {
		t.Error(err)
		return
	}
	for _, service := range services {
		fmt.Println(service.Service)
		fmt.Println(service.Address)
		fmt.Println(service.Port)
	}
}

func TestCfg_FindServiceById(t *testing.T) {
	n := NewServicesId("861848bf-755f-419f-badb-295dea458c6d", "10.143.24.84:8500")
	id, a, err := n.FindServiceById()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(id)
	fmt.Println(a)
}
