package activity

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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

const (
	follow = `{"@context":"https://www.w3.org/ns/activitystreams","id":"https://mstdn.social/04f2a2dc-a054-4d1a-b87d-447b0affecc3","type":"Follow","actor":"https://mstdn.social/users/hvturingga","object":"https://halfmemories.com/u/hvturingga"}`
	undo   = `{"@context":"https://www.w3.org/ns/activitystreams","id":"https://mstdn.social/users/hvturingga#follows/751608/undo","type":"Undo","actor":"https://mstdn.social/users/hvturingga","object":{"id":"https://mstdn.social/04f2a2dc-a054-4d1a-b87d-447b0affecc3","type":"Follow","actor":"https://mstdn.social/users/hvturingga","object":"https://halfmemories.com/u/hvturingga"}}`
)

func TestInboxActivity_Activity(t *testing.T) {
	fmt.Println(follow)
	fmt.Println(undo)
}
