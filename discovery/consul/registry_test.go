package consul

import (
	"testing"
)

const (
	consulAddr   = "hvxahv.disism.internal:8500"
	localhost    = "hvxahv.disism.internal"
	registerPort = "8080"
)

func TestCfg_Register(t *testing.T) {
	n, err := New("gateway",
		localhost,
		registerPort,
		[]string{"http", "gateway"},
		SetAddress(consulAddr),
	)
	if err != nil {
		t.Error(err)
	}

	if err := n.Register(); err != nil {
		t.Error(err)
		return
	}
}

func TestCfg_Deregister(t *testing.T) {
	n := NewServicesId("gateway", consulAddr)

	if err := n.Deregister(n.id); err != nil {
		t.Error(err)
		return
	}
}
