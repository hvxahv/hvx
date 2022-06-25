package internal

import (
	"context"
	"fmt"
	"testing"
)

func TestChannel_AddSubscriber(t *testing.T) {
	c := &channel{}
	subscriber, err := c.AddSubscriber(context.Background(), &v1alpha1.AddSubscriberRequest{
		ChannelId: "747232968277196801",
		AccountId: "746932029522116609",
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
		ChannelId: "747232968277196801",
		AccountId: "746932029522116609",
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
		AdminId:      "746931987134185473",
		ChannelId:    "747232969484730369",
		SubscriberId: "746932029522116609",
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(subscriber)
}

func TestChannel_GetAllSubscribers(t *testing.T) {
	c := &channel{}
	subscribers, err := c.GetAllSubscribers(context.Background(), &v1alpha1.GetAllSubscribersRequest{
		ChannelId: "747232969484730369",
		AdminId:   "746931987134185473",
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(subscribers.Subscriber)
}
