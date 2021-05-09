package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestInitConfig(t *testing.T) {
	if err := InitConfig(); err != nil {
		t.Log(err)
	}
	x := viper.GetString("port.accounts")
	fmt.Println(x)
}
