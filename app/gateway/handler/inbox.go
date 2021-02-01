package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/client/social"
	"hvxahv/pkg/structs"
	"hvxahv/pkg/utils"
	"log"
)

// InboxHandler 用于处理接收到收件箱的数据并调用 Inbox 客户端，接收客户端返回的数据
func InboxHandler(c *gin.Context) {
	name := c.Param("user")
	e, reqId, actor := GetEvent(c)
	data := &structs.Inbox{
		Name: name,
		RequestId: reqId,
		EventType:  e,
		Actor: actor,
	}
	r, err := social.InboxClient(data)
	if err != nil {
		log.Println(err)
	}
	log.Println("--------------------> INBOX 客户端返回的消息:", r)

	//data, _ := ioutil.ReadAll(c.Request.Body)
	//log.Printf("ctx.Request.body: %v", string(data))
}

// GetEvent 用于解析接收到的数据，并按照类型返回
func GetEvent(c *gin.Context) (string, string, string) {
	f := make(map[string]interface{})
	if err := c.BindJSON(&f); err != nil {
		log.Println(err)
	}
	// 进行接口断言取得 id 和 actor
	id, ok := f["id"].(string)
	if !ok {
		log.Println("获取 INBOX 事件 ID 断言失败得到的 id 不是字符串")
	}
	log.Println("进行断言之后得到的ID :",id)
	actor, ok := f["actor"].(string)
	if !ok {
		log.Println("获取 INBOX 事件 Actor 断言失败得到的 id 不是字符串")
	}

	log.Println("接收到的请求类型：", f["type"], "得到的事件 ID：", id, "用户：", f["actor"])
	switch f["type"] {
	case "Follow":
		return "follow", id, actor
	case "Undo":
		return "undo", id, actor
	default:
		return "", "", ""
	}
}

// GetInboxHandler
func GetInboxHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	r, err := social.GetInboxClient(name)
	if err != nil {
		log.Println(err)
	}
	log.Println("--------------------> 用户 GET INBOX 客户端返回的消息:", r)
}