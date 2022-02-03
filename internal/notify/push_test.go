package notify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"log"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/push"
)

func TestNotify_Push(t *testing.T) {
	data, err := json.Marshal(push.NewData("Notify",
		"Life's Not Out To Get You!",
		"https://avatars.githubusercontent.com/u/94792300?s=200&v=4",
		"Authorized"))
	if err != nil {
		log.Println(err)
		return
	}
	d := &v1alpha1.NewNotifyPush{
		DeviceId: "732570178888761345",
		Data:     data,
	}
	s := notify{}
	reply, err := s.Push(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(reply.Code, reply.Reply)
}
