package activity

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func IniTestConfig(t *testing.T) {
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

	// If a configs file is found, read it in.
	if err3 := viper.ReadInConfig(); err3 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	n := cockroach.NewDBAddr()
	if err2 := n.InitDB(); err2 != nil {
		return
	}
}

func TestNewMessages(t *testing.T) {
	IniTestConfig(t)
	//nm := NewInbox("https://mas.to/users/hvturingga", "Follow", "https://mas.to/9dc74aa1-9f24-4514-b6e4-00ebd6be012f", "hvturingga")
	//nm.NewInbox()
}

func TestInboxes_GetInboxesByObjectID(t *testing.T) {
	IniTestConfig(t)

	r, err := NewObjectID(698619813575491585).GetInboxesByID()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}