package cfg

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	defaultConfigName = ".hvxahv"
)

func DefaultConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".account" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(defaultConfigName)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if err := cockroach.NewRoach().Dial(); err != nil {
		fmt.Println(err)
		return
	}
}
