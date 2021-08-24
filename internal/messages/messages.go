package messages

import (
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/pkg/remote"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

//{
//	"@context":"https://www.w3.org/ns/activitystreams",
//	"id":"https://mas.to/users/hvturingga#follows/113972/undo",
//	"type":"Undo",
//	"actor":"https://mas.to/users/hvturingga",
//	"object":{
//		"id":"https://mas.to/30ff54b1-c2dd-482c-ad70-43a775476584",
//		"type":"Follow","actor":"https://mas.to/users/hvturingga",
//		"object":"https://07ee-2408-832f-20b2-be60-7c3c-bb0d-7b8b-bb20.ngrok.io/u/hvturingga"
//	}
//}

type Reply struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Actor  string `json:"actor"`
		Object string `json:"object"`
	} `json:"object"`
}

//{
//	"@context":"https://www.w3.org/ns/activitystreams",
//	"id":"https://mas.to/e27a4e0e-a0a0-400e-a395-6b0e60f08291",
//	"type":"Follow",
//	"actor":"https://mas.to/users/hvturingga",
//	"object":"https://07ee-2408-832f-20b2-be60-7c3c-bb0d-7b8b-bb20.ngrok.io/u/hvturingga"
//}

type Messages struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

func NewMessages(types string, actor string, object string) *Messages {
	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id = fmt.Sprintf("https://%s/%s", viper.GetString("localhost"), uuid.New().String())
		a = fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), actor)
	)
	fmt.Println("请求的 ID :", id)
	fmt.Println("当前用户的地址:", a)
	return &Messages{Context: ctx, Id: id, Type: types, Actor: a, Object: object}
}

func (i *Messages) Outbox() {
	m := &Messages{
		Context: i.Context,
		Id:      i.Id,
		Type:    i.Type,
		Actor:   i.Actor,
		Object:  i.Object,
	}
	data, err := json.Marshal(*m)
	if err != nil {
		return 
	}
	remote.SendRequest(i.Object, data)
}

func (i *Messages) Inbox(name string) {
	fmt.Printf("%s 给 %s 发送了请求", i.Actor, name)
	switch i.Type {
	case "Follow":
		fmt.Printf("请求关注")
	case "Undo":
		fmt.Printf("取消了请求")

	}
}

type Message interface {
	Inbox(string)
	Outbox()
}
