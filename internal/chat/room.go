package chat

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/matrix"
	"github.com/matrix-org/gomatrix"
)

type Auth struct {
	Username string
	Token    string
}

type ReqCreateRoom struct {
	Auth *Auth
	Data *gomatrix.ReqCreateRoom
}

func NewReqCreateRoom(username, token string, data *gomatrix.ReqCreateRoom) *ReqCreateRoom {
	return &ReqCreateRoom{
		Auth: &Auth{
			Username: username,
			Token:    token,
		},
		Data: data,
	}
}

func (r *ReqCreateRoom) CreateRoom() error {
	c, err2 := matrix.NewClient(r.Auth.Username, r.Auth.Token)
	if err2 != nil {
		return err2
	}

	room, err := c.CreateRoom(r.Data)
	if err != nil {
		return err
	}
	fmt.Println(room)
	return nil

}

type Matrix interface {
	CreateRoom() error
}
