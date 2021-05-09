package config

import (
	"fmt"
	"github.com/spf13/viper"
)
const (
	fileAddr = "../../configs/config.yaml"
	remoteAddr = ""
)

// InitConfig Config Get Config file
func InitConfig(flag string) error {
	if flag == "local" {
		viper.SetConfigFile(fileAddr)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
			return err
		}
		return nil
	} else {

		// consul kv
		return nil
	}
}
