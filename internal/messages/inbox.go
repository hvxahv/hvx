package messages

import (
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/activitypub"
	"log"
)

type Inbox struct {
	Context string             `json:"@context"`
	Id      string             `json:"id"`
	Type    string             `json:"type"`
	Actor   string             `json:"actor"`
	Object  activitypub.Object `json:"object"`
}

func (i *Inbox) Inbox(name string) {
	// TODO - INBOX DATA
	fmt.Printf("%s 给 %s 发送了请求; 请求 ID : %s", i.Actor, name, i.Id)
	switch i.Type {
	case "Follow":
		fmt.Println("请求关注你")
		// Check if there is this actor
		nm := NewMessages(i.Actor, i.Type, i.Id, name)
		nm.New()

	case "Undo":
		fmt.Printf("取消了请求")
		fmt.Println("得到的接口数据:", i.Object)
		nm := NewMessages(i.Actor, i.Type, i.Id, name)
		nm.New()

	case "Reject":
		fmt.Printf("拒绝了你的请求")
		fmt.Println("接收了你的请求:", i.Object)
		nm := NewMessages(i.Actor, i.Type, i.Id, name)
		nm.New()

	case "Accept":
		fmt.Println("接收了你的请求:", i.Object)
		nm := NewMessages(i.Actor, i.Type, i.Id, name)
		nm.New()

		nf := accounts.NewFollow(i.Actor, name)
		err := nf.New()
		if err != nil {
			log.Println(err)
		}
	}
}
