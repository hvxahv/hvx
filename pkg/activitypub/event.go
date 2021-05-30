package activitypub

import (
	"github.com/gin-gonic/gin"
	"log"
)

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
	actor, ok := f["actor"].(string)
	if !ok {
		log.Println("获取 INBOX 事件 Actor 断言失败得到的 id 不是字符串")
	}

	// 通过事件 ID 可以查看 Activitypub 的元数据
	log.Println("-------------------------------------")
	log.Println("接收到的请求类型：", f["type"])
	log.Println("得到的活动事件 ID：", id)
	log.Println("获取活动用户：", f["actor"])
	log.Println("-------------------------------------")

	switch f["type"] {
	case "Follow":
		return "follow", id, actor
	case "Undo":
		return "undo", id, actor
	case "Create":
		return "Create", id, actor
	case "Accept":
		return "Accept", id, actor

	default:
		return "", "", ""
	}
}
