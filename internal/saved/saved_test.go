package saved

import (
	"fmt"
	"log"
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
	if err2 := viper.ReadInConfig(); err2 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}
	// Initialize the database.
	if err := cockroach.NewDBAddr().InitDB(); err != nil {
		fmt.Println(err)
		return
	}
}

func TestSaved_Create(t *testing.T) {
	if err := NewSaves(123123, "xs", "hash", "text").Create(); err != nil {
		log.Println(err)
		return
	}
}
