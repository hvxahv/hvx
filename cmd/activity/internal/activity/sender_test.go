package activity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/activitypub"
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

func TestSender_Do(t *testing.T) {
	const (
		actorAddress  = "https://halfmemories.com/u/hvturingga"
		inbox         = "https://mastodon.disism.com/users/hvturingga/inbox"
		object        = "https://mastodon.disism.com/users/hvturingga"
		pemPrivateKey = ``
	)
	body := &activitypub.Follow{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      fmt.Sprintf("%s/%s", actorAddress, uuid.NewString()),
		Type:    "Follow",
		Actor:   actorAddress,
		Object:  object,
	}
	marshal, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(marshal)
}
