package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/client/account"
	"hvxahv/pkg/utils"
	"log"
)

// NewAccountsHandler ... 创建账户的 Handler
func NewAccountsHandler(c *gin.Context) {
	u := c.PostForm("username")
	p := c.PostForm("password")
	// 将得到的数据交给客户端访问服务
	r, err := account.NewAccountClient(u, p)
	if err != nil {
		log.Println(err)
	} else {
		account.NewAccountsResponse(c, r)
	}
}

// GetAccountsHandler ... 获取账户信息的 Handler
// 通过解析 Token 得到的用户名获取账户信息
// 不同于 Actor ，因为它是给用户使用的接口
func GetAccountsHandler(c *gin.Context) {
	author := utils.GetUserName(c)
	r, err := account.GetAccountsClient(author)
	if err != nil {
		log.Println(err)
	} else {
		account.AccountsResponse(c, r)
	}
}

// GetActorHandler ... 通过其他实例的请求 Actor 来获取账户信息
// 不同于 Accounts ，因为它是给服务器使用的接口，详情请查阅 activitypub actor 获取相关信息
func GetActorHandler(c *gin.Context) {
	name := c.Param("user")
	r, err := account.GetActorClient(name)
	if err != nil {
		log.Println(err)
	} else {
		account.ActorResponse(c, r)
	}
}
// GetWebFingerHandler
// 给查询的服务器返回 WebFinger，常与 Actor 一起使用
// 接受 resource 参数并交给客户端进一步处理
func GetWebFingerHandler(c *gin.Context) {
	name := c.Query("resource")

	r, err := account.GetWebFingerClient(name)
	if err != nil {
		log.Println(err)
	}
	account.WebFingerResponse(c, r)
}

func DeleteAccountHandler(c *gin.Context) {

}

func AccountSettingHandler(c *gin.Context)  {

}

// GetActorOutBox 获取用户的 outbox
func GetActorOutbox(c *gin.Context) {
	name := c.Param("user")
	log.Println(name, "用户请求了数据")

	account.OutboxResponse(c)
}

