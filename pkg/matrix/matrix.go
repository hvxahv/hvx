/*
Matrix is an open standard for interoperable, decentralised, real-time communication over IP.
https://matrix.org/docs/guides/introduction
*/

package matrix

import (
	"github.com/matrix-org/gomatrix"
	"github.com/spf13/viper"
)

func NewClient(username, token string) (*gomatrix.Client, error) {
	addr := viper.GetString("matrix.addr")
	cli, err := gomatrix.NewClient(addr, username, token)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
