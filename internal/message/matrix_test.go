package message

import (
	"fmt"
	"github.com/hvxahv/hvx/pkg/cockroach"
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

func TestMatrices_Create(t *testing.T) {
	a := NewMatrixAccesses(733124680636596225, "xxs", "localhost", "id", "did")
	if err := a.Create(); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestMatrices_UpdateToken(t *testing.T) {
	a := NewAccessUpdateToken(733124680636596225, "xvJGvzoZLBSvt8oKbGHZkuTNIrshqhFHjZkA5Rt-vjox")
	if err := a.UpdateToken(); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}
