package db

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".accounts" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".accounts")

	viper.AutomaticEnv() // read in environment variables that match

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	nd := NewDb()
	if err := nd.InitDB(); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}

func TestCreateDB(t *testing.T) {

	nd := NewDb()
	name := "hvxahv"
	if err := nd.Create(name); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}
