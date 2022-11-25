package cfg

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigName     = ".hvx"
	defaultConsulEndpoint = "127.0.0.1:8500"
)

func HomeConf() error {

	// Find home directory.
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(defaultConfigName)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		return errors.New(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	}
	return nil
}

func LocalDefault() error {
	viper.SetConfigName(defaultConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../conf")

	if err := viper.ReadInConfig(); err != nil {
		return errors.New(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	}
	return nil
}

func Default() {

	if err := LocalDefault(); err != nil {
		r := New(ConsulProvider, defaultConsulEndpoint, defaultConfigName)
		if err := r.Dial(); err != nil {
			fmt.Println(err)
			return
		}
	}

	if err := cockroach.NewRoach().Dial(); err != nil {
		fmt.Println(err)
		return
	}
}
