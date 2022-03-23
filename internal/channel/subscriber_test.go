package channel

import (
	"context"
	"fmt"
	"github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"testing"
)

func TestChannel_AddSubscriber(t *testing.T) {
	c := &channel{}
	subscriber, err := c.AddSubscriber(context.Background(), &v1alpha1.AddSubscriberRequest{
		ChannelId: "746637380461068289",
		AccountId: "7469260235139645451",
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(subscriber)
}

func TestChannel_Unsubscribe(t *testing.T) {
	c := &channel{}
	unsubscribe, err := c.Unsubscribe(context.Background(), &v1alpha1.UnsubscribeRequest{
		ChannelId: "746637380461068289",
		AccountId: "746926023513964545",
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(unsubscribe)
}

func TestChannel_RemoveSubscriber(t *testing.T) {
	c := &channel{}
	subscriber, err := c.RemoveSubscriber(context.Background(), &v1alpha1.RemoveSubscriberRequest{
		AdminId:      "746926023513964545",
		ChannelId:    "746637380461068289",
		SubscriberId: "7469260235139645451",
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(subscriber)
}
