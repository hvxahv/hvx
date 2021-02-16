/*
	Gateway 用户验证
*/
package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hvxahv/pkg/client/accounts"
	"hvxahv/pkg/models"
	"hvxahv/pkg/response"
	"hvxahv/pkg/utils"
	"log"
)

// VerificationHandler 登录时的用户验证与 Token 发放
func VerificationHandler(c *gin.Context) {
	u := c.PostForm("username")
	p := c.PostForm("password")

	a := models.Accounts{ Username: u, Password: p }

	// 调用验证客户端获取返回的账户信息
	r, err := accounts.VerifyAccountsClient(u)
	if err != nil {
		log.Println("进行账户登录验证时出现错误: ", err)
	}
	if r.Username == "" {
		response.SimpleResponse(c, "202", "用户不存在")
		return
	}

	// 登录密码验证，密码正确将生成 token 返回
	if err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(p)); err != nil {
		response.SimpleResponse(c, "401", "密码错误")
		return
	}
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
