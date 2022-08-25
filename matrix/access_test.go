package matrix

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
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
	deviceId := uuid.New().String()
	register, err := NewRegisterReq(deviceId, "idinahui8964", "hvxahv123").Register()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(register.Body))
}

func TestDeactivateReq_DeactivateReq(t *testing.T) {
	req, err := NewDeactivateReq("").DeactivateReq()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(req)
}

func TestNewEditPasswordReq(t *testing.T) {
	password, err := NewEditPasswordReq("", true).EditPassword()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(password)
}
