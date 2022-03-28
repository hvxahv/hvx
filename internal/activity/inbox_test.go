package activity

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/activity/v1alpha1"
	"golang.org/x/net/context"
	"testing"
)

func TestActivity_InboxForFollow(t *testing.T) {
	name := "hvxahv"
	follow := []byte(`{
		"@context":"https://www.w3.org/ns/activitystreams",
		"id":"https://mas.to/e27a4e0e-a0a0-400e-a395-6b0e60f08291",
		"type":"Follow",
		"actor":"https://mas.to/users/hvturingga",
		"object":"https://07ee-2408-832f-20b2-be60-7c3c-bb0d-7b8b-bb20.ngrok.io/u/hvturingga"
	}`)

	s := &activity{}
	inbox, err := s.Inbox(context.Background(), &pb.InboxRequest{
		Name: name,
		Data: follow,
	})
	if err != nil {
		return
	}
	fmt.Println(inbox)
}

func TestActivity_InboxForUndo(t *testing.T) {
	name := "hvxahv"
	undo := []byte(`{
	  "@context": "https://www.w3.org/ns/activitystreams",
	  "summary": "Sally retracted her offer to John",
	  "type": "Undo",
	  "actor": "https://mas.to/users/hvturingga",
	  "object": {
		"type": "Follow",
		"actor": "https://halfmemories.com/u/hvturingga",
		"object": "https://mas.to/e27a4e0e-a0a0-400e-a395-6b0e60f08291"
	  }
	}`)
	s := &activity{}
	inbox, err := s.Inbox(context.Background(), &pb.InboxRequest{
		Name: name,
		Data: undo,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(inbox)
}

func TestActivity_InboxForAccept(t *testing.T) {
	name := "hvxahv"
	accept := []byte(`{
	  "@context": "https://www.w3.org/ns/activitystreams",
	  "summary": "Sally accepted an invitation to a party",
	  "type": "Accept",
	  "actor": "https://mas.to/users/hvturingga",
	  "object": {
		"type": "Invite",
		"actor": "https://halfmemories.com/u/hvturingga",
		"object": "https://mas.to/e27a4e0e-a0a0-400e-a395-6b0e60f08291"
	  }
	}`)
	s := &activity{}
	inbox, err := s.Inbox(context.Background(), &pb.InboxRequest{
		Name: name,
		Data: accept,
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
		ActivityId: "748289222292602881",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(inbox)
}

func TestActivity_GetInboxesByAccountID(t *testing.T) {
	s := &activity{}
	inboxes, err := s.GetInboxesByAccountID(context.Background(), &pb.GetInboxesByAccountIDRequest{
		AccountId: "746931986864701441",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(inboxes)
}

func TestActivity_DeleteInboxByActivityID(t *testing.T) {
	s := &activity{}
	ibx, err := s.DeleteInboxByInboxesID(context.Background(), &pb.DeleteInboxByInboxesIDRequest{
		ActivityId: "748289253916311553",
		AccountId:  "746931986864701441",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ibx)
}
