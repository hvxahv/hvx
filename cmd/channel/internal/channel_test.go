package internal

import (
	"fmt"
	"os"
	"testing"

	"github.com/hvxahv/hvx/cockroach"

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
	if err := cockroach.NewRoach().Dial(); err != nil {
		fmt.Println(err)
		return
	}
}

func TestChannel_CreateChannel(t *testing.T) {
	const (
		actorId   = 786085616327884801
		creatorId = 785518573776797697
	)
	if err := NewChannels(actorId, creatorId, "PUBLIC_KEY").CreateChannel(); err != nil {
		t.Error(err)
	}
}

func TestChannel_GetChannelsByAccountID(t *testing.T) {
	channels, err := NewChannelsCreatorId(786085616327884801).GetChannels()
	if err != nil {
		t.Error(err)
	}
	t.Log(channels)
}

func TestChannel_DeleteChannel(t *testing.T) {
	const (
		channelId = 786085616327884801
		creatorId = 785518573776797697
	)
	if err := NewChannelsDelete(channelId, creatorId).DeleteChannel(); err != nil {
		t.Error(err)
	}
}

func TestChannel_DeleteAllChannelsByAccountID(t *testing.T) {
	const (
		channelId = 786085616327884801
		creatorId = 785518573776797697
	)
	if err := NewChannelsDelete(channelId, channelId).DeleteChannels(); err != nil {
		t.Error(err)
	}
}
