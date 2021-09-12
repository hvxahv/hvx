package channel

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cache"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/ipfs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInitChannelConfig(t *testing.T) {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err1 := viper.ReadInConfig(); err1 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	n := cockroach.NewDBAddr()
	if err2 := n.InitDB(); err2 != nil {
		return
	}

	// If a configs file is found, read it in.
	if err3 := viper.ReadInConfig(); err3 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	ipfs.InitIPFS()

	cache.InitRedis(1)

}

func TestNewChannels(t *testing.T) {
	TestInitChannelConfig(t)

	nc1 := NewChannels("Hvxahv", "", "avatar", "bio", "hvturingga", false)
	fmt.Println(nc1)
	err := nc1.New()
	if err != nil {
		return
	}

	//nc1 := NewChannels("disism", "", "avatar", "bio", "hvturingga", false)
	//fmt.Println(nc1)
	//err := nc1.New()
	//if err != nil {
	//	return
	//}
}

func TestChannels_FetchByLink(t *testing.T) {
	TestInitChannelConfig(t)

	n := NewChannelsByLink("y2yrpHc0TTK-4AB")
	data, err := n.QueryByLink()
	if err != nil {
		return
	}

	fmt.Println(data.Name)
}
