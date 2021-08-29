package messages

import (
	"fmt"
)

func (i *Inbox) Inbox(name string) {

	// TODO - INBOX DATA
	fmt.Printf("%s 给 %s 发送了请求", i.Actor, name)
	switch i.Type {
	case "Follow":
		fmt.Printf("请求关注")
	case "Undo":
		fmt.Printf("取消了请求")
		fmt.Println("得到的接口数据:", i.Object)
	case "Reject":
		fmt.Printf("拒绝了你的请求")
	}
}


