package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/app/gateway/client/account"
	"hvxahv/pkg/auth"
	"log"
)

// NewAccountsHandler ... 创建账户的 Handler
func NewAccountsHandler(c *gin.Context) {
	u := c.PostForm("username")
	p := c.PostForm("password")

	r, err := account.NewAccountClient(u, p)
	if err != nil {
		log.Println(err)
	} else {
		account.AccountsHandlerResponse(r, c)
	}
}

// GetAccountsHandler ...
func GetAccountsHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	account.GetAccountsClient(author)
}

func DeleteAccountHandler(c *gin.Context) {

}

func AccountSettingHandler(c *gin.Context)  {

}

