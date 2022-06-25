package internal

import (
	"context"
	"fmt"

	"testing"
)

func TestChannel_CreateBroadcast(t *testing.T) {
	c := &channel{}
	broadcast, err := c.CreateBroadcast(context.Background(), &pb.CreateBroadcastRequest{
		ChannelId: "747232969484730369",
		AdminId:   "746932029522116609",
		ArticleId: "1234",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(broadcast.Reply)
}

func TestChannel_GetAllBroadcasts(t *testing.T) {
	c := &channel{}
	broadcasts, err := c.GetAllBroadcasts(context.Background(), &pb.GetAllBroadcastsRequest{
		ChannelId: "747232969484730369",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(broadcasts.Broadcasts)
}

func TestChannel_DeleteBroadcast(t *testing.T) {
	c := &channel{}
	ch, err := c.DeleteBroadcast(context.Background(), &pb.DeleteBroadcastRequest{
		ChannelId:   "747232969484730369",
		AdminId:     "746932029522116609",
		BroadcastId: "747492082502762497",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ch.Reply)
}
