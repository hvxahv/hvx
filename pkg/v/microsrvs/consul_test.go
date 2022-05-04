package microsrvs

import (
	"fmt"
	"github.com/hvxahv/hvx/pkg/cache"
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

func TestRegister_Register(t *testing.T) {
	p := viper.GetString("x.gateway.bac.port")

	tags := []string{"gateway.bac", "http", "RESTFul"}
	nr := NewRegister("gateway.bac", p, tags, "localhost")
	err := nr.Register()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func TestDeregister(t *testing.T) {

	err := Deregister("b5e6670f-33e9-4d92-a186-583c64aa43ce")
	if err != nil {
		fmt.Println(err)
	}
}
