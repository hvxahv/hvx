/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package ipfs

import (
	ipfs "github.com/ipfs/go-ipfs-api"
)

type is struct {
	shell *ipfs.Shell
}

//
//func InitIPFS() {
//	sh := ipfs.NewShell(viper.GetString("ipfs_addr"))
//	is = sh
//}
//
//func GetIPFS() *ipfs.Shell {
//	return is
//}
