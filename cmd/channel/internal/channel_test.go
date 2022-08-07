package internal

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func TestChannel_CreateChannel(t *testing.T) {

}

func TestChannel_GetChannelsByAccountID(t *testing.T) {

}

func TestChannel_DeleteChannel(t *testing.T) {

}

func TestChannel_DeleteAllChannelsByAccountID(t *testing.T) {

}
