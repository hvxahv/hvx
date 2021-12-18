package notify

import (
	"fmt"
	"os"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/cockroach"
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
	n := cockroach.NewDBAddr()
	if err := n.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
		return
	}

}

func TestNotify_Create(t *testing.T) {
	if err := NewNotifies(720140086112714753,
		"https://sg2p.notify.windows.com/w/?token=BQYAAACwPljTW1WbNsGHH9jVFNzy3o73MaV%2bw1Im3ZKSrSQDIdJ0l%2fRMKTZ%2bJUhT44S5tncAr9VVyVBqIHpQqDFNtAEH1xjeQp1%2bjWcwQI%2bHfU%2bwcNdspTxPBKTaBiUlZVxWPK1XmslyuHXvxqGqLQqLX6eXl7vAhWrtVGk0A6z7agvQOY9d%2bmNMKakkvpPJB2rXhemdP%2fZYqc4SmEQXDjiJdN7Pg%2fWJQai8CB2gtzr%2bb1SuaQmnve%2f7SWIqN9ihHvI9oo7wagvULBS%2fAuC9sJ1IamUO45wQqbG%2f3oczBzfpjSuBM42bJ4MGqkN35ePU5rnXfYU%3d",
		"BI5N3A_zbQPVP6W6Wo-IzgOWu1ux0FsGHAxCBf-X0grLANXfgPd5r4tOCDT7o7zHvabkLItPs2evXrd3AHB4OXY",
		"OcHJ2-7QrDAm_Mp95TW6Hw",
	).Create(); err != nil {
		fmt.Println(err)
	}
}

func TestNotify_Get(t *testing.T) {
	r, err := NewNotifiesByDeviceID(720140086112714753).Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Auth)
}
