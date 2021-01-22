package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/app/ingress/client/account"
	"hvxahv/pkg/auth"
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
		// 将返回的状态码进行处理并返回给前端
		account.AccountsHandlerResponse(r, c)
	}
}

// GetAccountsHandler ...
func GetAccountsHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	r, err := account.GetAccountsClient(author)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"username": r.Username,
	})
}

func DeleteAccountHandler(c *gin.Context) {

}

func AccountSettingHandler(c *gin.Context)  {

}

