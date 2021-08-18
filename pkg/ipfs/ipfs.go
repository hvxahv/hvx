package ipfs

import (
	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
)
var is *ipfs.Shell

func InitIPFS()  {
	sh := ipfs.NewShell(viper.GetString("ipfs_addr"))
	is = sh
}

func GetIPFS() *ipfs.Shell {
	return is
}

