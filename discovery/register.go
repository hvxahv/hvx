package discovery

import "github.com/hvxahv/hvx/discovery/consul"

type Registrar interface {
	consul.Registry
}
