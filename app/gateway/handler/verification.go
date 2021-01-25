/*
	Gateway 用户验证
*/
package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hvxahv/app/gateway/client/account"
	"hvxahv/pkg/structs"
	"hvxahv/pkg/utils"
)

// VerificationHandler 登录时的用户验证与 Token 发放
func VerificationHandler(c *gin.Context) {
	u := c.PostForm("username")
	p := c.PostForm("password")

	a := structs.Accounts{ Username: u, Password: p }

	r, err := account.GetAccountsClient(u)
	if err != nil {
		c.JSON(202, gin.H{
			"state": "202",
			"message": "用户不存在",
		})
	}
	// 登录密码验证，密码正确将生成 token 返回
	if err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(p)); err != nil {
		c.JSON(401, gin.H{
			"state": "401",
			"message": "密码错误",
		})
	} else {
		t, err := utils.GenerateToken(a)
		if err != nil {
			fmt.Println("生成 Token 失败！")
		}
		c.JSON(200, gin.H{
			"state": "200",
			"message": "登陆成功，欢迎您：" + u,
			"token": t,
		})
	}
}

