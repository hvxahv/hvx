package internal

import (
	"github.com/hvxahv/hvx/cfg"
	"testing"
)

func init() {
	cfg.Default()
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
