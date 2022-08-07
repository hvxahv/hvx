package ipfs

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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
	// Initialize the database.
	if err := cockroach.NewRoach().Dial(); err != nil {
		fmt.Println(err)
		return
	}
}

func TestAddr_Add(t *testing.T) {
	data := strings.NewReader("disism.com")
	add, err := NewAddr().Add(data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(add)
}

func TestAddr_GetGatewayAddress(t *testing.T) {
	a := NewAddr().GetGatewayAddress("QmZVgg8gzfLaENL888YSd4DKt9MDkqT5h9KbxNLDmjW1aM")
	t.Log(a)
}
