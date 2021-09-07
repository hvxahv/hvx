package messages

import (
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/activitypub"
	"log"
)

type Inbox struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

func InboxEventHandler(name string, body []byte) {
	i := Inbox{}
	err := json.Unmarshal(body, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	fmt.Printf("%s 给 %s 发送了请求; 请求 ID : %s", i.Actor, name, i.Id)

	switch i.Type {
	case "Follow":
		f := activitypub.Follow{}
		err2 := json.Unmarshal(body, &f)
		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println(f)
		fmt.Printf("%s 请求关注你", f.Actor)

		nm := NewMessages(f.Actor, f.Type, f.Id, name)
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
		fmt.Println("接受了你的请求:", i.Object)
		a := activitypub.Accept{}
		err2 := json.Unmarshal(body, &a)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(a)

		nm := NewMessages(a.Actor, a.Type, a.Id, name)
		nm.New()
		
		nf := accounts.NewFollow(name, i.Actor)
		err3 := nf.New()
		if err3 != nil {
			log.Println(err3)
		}
	}
}
func (i *Inbox) Inbox(name string) {

}
