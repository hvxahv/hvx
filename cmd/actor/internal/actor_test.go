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

func TestActors_IsExist(t *testing.T) {
	a := NewActorsIsExist("example.com", "alice")
	if a.IsExist() {
		t.Log("Actor exists")
	} else {
		t.Log("Actor does not exist")
	}
}

func TestActors_Create(t *testing.T) {

}

func TestActors_Get(t *testing.T) {

}

func TestActors_GetActorsByPreferredUsername(t *testing.T) {

}

func TestActors_Add(t *testing.T) {

}

func TestGetActorByUsername(t *testing.T) {

}

func TestActors_Edit(t *testing.T) {
	a := NewActorsId(1)
	if err := a.Edit(); err != nil {
		t.Error(err)
	}
}

func TestActors_Delete(t *testing.T) {
	a := NewActorsId(1)
	if err := a.Delete(); err != nil {
		t.Error(err)
	}
}
