package ipfs

import (
	"fmt"
	"io"

	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
)

type addr struct {
	gateway string
	ipfs    *ipfs.Shell
}

type IPFS interface {
	Add(data io.Reader) (cid string, err error)
	GetGatewayAddress(cid string) (address string)
}

func NewAddr() *addr {
	address := viper.GetString("ipfs.address")
	sh := ipfs.NewShell(address)
	return &addr{
		gateway: viper.GetString("ipfs.gateway"),
		ipfs:    sh,
	}
}

func (a *addr) Add(data io.Reader) (cid string, err error) {
	add, err := a.ipfs.Add(data)
	if err != nil {
		return "", err
	}
	return add, nil
}

func (a *addr) GetGatewayAddress(cid string) (address string) {

	return fmt.Sprintf("%s%s", a.gateway, cid)
}
