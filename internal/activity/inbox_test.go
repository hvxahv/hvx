package activity

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/activity/v1alpha1"
	"golang.org/x/net/context"
	"testing"
)

func TestActivity_CreateInbox(t *testing.T) {
	name := "hvxahv"
	follow := []byte(`{
	"@context":"https://www.w3.org/ns/activitystreams",
	"id":"https://mas.to/e27a4e0e-a0a0-400e-a395-6b0e60f08291",
	"type":"Follow",
	"actor":"https://mas.to/users/hvturingga",
	"object":"https://07ee-2408-832f-20b2-be60-7c3c-bb0d-7b8b-bb20.ngrok.io/u/hvturingga"
}`)

	s := &activity{}
	inbox, err := s.CreateInbox(context.Background(), &pb.CreateInboxRequest{
		Name: name,
		Data: follow,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(inbox)
}

func TestActivity_GetInboxByActivityID(t *testing.T) {
	s := &activity{}
	inbox, err := s.GetInboxByActivityID(context.Background(), &pb.GetInboxByActivityIDRequest{
		ActivityId: "747526412430999553",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(inbox)
}
