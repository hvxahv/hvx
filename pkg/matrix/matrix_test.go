package matrix

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/cache"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func init() {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	cache.InitRedis(1)

}

func TestNewClient(t *testing.T) {
	c, err := NewClient("", "")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(c)
}
