package chat

import (
	"fmt"
	"github.com/matrix-org/gomatrix"
	"testing"
)

func TestReqCreateRoom_CreateRoom(t *testing.T) {
	TestInitDB(t)
	data := &gomatrix.ReqCreateRoom{
		Visibility:      "",
		RoomAliasName:   "hvxahv-dev",
		Name:            "HVXAHV_DEV",
		Topic:           "",
		Invite:          nil,
		Invite3PID:      nil,
		CreationContent: nil,
		InitialState:    nil,
		Preset:          "",
		IsDirect:        false,
	}
	n := NewReqCreateRoom("hvturingga", "au7l_K7D58mXztY3_UXKR8KVOFwY0fH3oIiubAJswZU", data)
	if err := n.CreateRoom(); err != nil {
		fmt.Println(err)
	}

}
