package internal

import (
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"testing"
)

func TestSubscribes_GetSubscribers(t *testing.T) {
	ctx := context.Background()
	subscribers, err := clientv1.New(ctx, microsvc.ChannelServiceName).GetSubscribers(807240711882047489, 801935105807482881)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(subscribers)
}
