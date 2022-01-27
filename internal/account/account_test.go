package account

import (
	"fmt"
	"github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"golang.org/x/net/context"
	"os"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/cache"
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

	cache.InitRedis(1)
}

func TestAccount_Create(t *testing.T) {
	d := &v1alpha1.NewAccountCreate{
		Username:  "hvturingga",
		Mail:      "hvturingga@disism.com",
		Password:  "hvxahv123",
		PublicKey: "p",
	}
	s := &account{}
	create, err := s.Create(context.Background(), d)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(create)
}
