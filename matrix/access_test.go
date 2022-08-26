package matrix

import (
	"fmt"
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".sok" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
func TestRegisterReq_Register(t *testing.T) {

}

func TestDeactivateReq_DeactivateReq(t *testing.T) {

}

func TestNewEditPasswordReq(t *testing.T) {

}
