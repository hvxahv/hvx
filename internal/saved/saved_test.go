package saved

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/ipfs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err2 := viper.ReadInConfig(); err2 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}
}

func TestUpload(t *testing.T) {
	TestInitConfig(t)

	ipfs.InitIPFS()

	Upload()
}

