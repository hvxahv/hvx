package message

import (
	"fmt"
	"os"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestInitChatConfig(t *testing.T) {

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

}

func TestAccess_Register(t *testing.T) {
	TestInitChatConfig(t)
}

func TestMatrix_Create(t *testing.T) {
	TestInitChatConfig(t)

	c := NewMatrixAccesses(111, "", "", "", "")
	if err := c.Create(); err != nil {
		fmt.Println(err)
	}
}
