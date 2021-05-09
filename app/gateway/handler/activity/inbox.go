package activity

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/server/middleware"
	social2 "hvxahv/internal/client/social"
	"hvxahv/internal/inbox"
	"hvxahv/pkg/activitypub"
	"log"
)

/**
InboxHandler 这是一个公共的收件箱
用于处理接收到收件箱的数据并调用 Inbox 客户端，接收客户端返回的数据
 */
func InboxHandler(c *gin.Context) {
	name := c.Param("user")
	e, reqId, actor := activitypub.GetEvent(c)
	data := &inbox.Inbox{
		Name: name,
		RequestId: reqId,
		EventType:  e,
		Actor: actor,
	}
	r, err := social2.InboxClient(data)
	if err != nil {
		log.Println(err)
	}
	// TODO 将客户端返回的消息同通知给用户
	log.Println("--------------------> INBOX 客户端返回的消息:", r)

	//data, _ := ioutil.ReadAll(c.Request.Body)
	//log.Printf("ctx.Request.body: %v", string(data))
}

// GetInboxHandler 获取用户的收件箱内容
func GetInboxHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	r, err := social2.GetInboxClient(name)
	if err != nil {
		log.Println(err)
	}
	log.Println("--------------------> 用户 GET INBOX 客户端返回的消息:", r)
	c.JSON(200, r)
}

func saveFollowing() {

}