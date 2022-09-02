package internal

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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

func TestFs_Create(t *testing.T) {
	if err := NewFsCreate(793098623322324993,
		"aa8ee033-761e-4069-a976-37801ea5ba7f-mmexport1661844042932.jpg",
		"",
	).Create(); err != nil {
		t.Error(err)
		return
	}
}

func TestFs_Delete(t *testing.T) {
	if err := NewFs(793098623322324993,
		"aa8ee033-761e-4069-a976-37801ea5ba7f-mmexport1661844042932.jpg",
	).Delete(); err != nil {
		t.Error(err)
		return
	}
}

func TestFs_Get(t *testing.T) {
	fd, err := NewFs(793098623322324993,
		"233b4726-1041-4def-8a46-bbb60f8da121-mmexport1661844037844.jpg",
	).Get()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fd)
}
