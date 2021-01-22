/*
	Ingress Gateway 用户验证实现
*/
package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hvxahv/app/accounts/app"
	"hvxahv/pkg/auth"
	"hvxahv/pkg/database"
	"log"
)

func VerificationHandler(c *gin.Context) {
	u := c.PostForm("username")
	p := c.PostForm("password")

	a := app.Accounts{ Username: u, Password: p }

	db, err := database.NewDB()
	if err != nil {
		log.Println("Database Connect Error", err)
	}
	if db.Debug().Table("accounts").Where("username = ?", u).First(&a).RecordNotFound() {
		c.JSON(202, gin.H{
			"state": "202",
			"message": "用户不存在",
		})
	} else {
		// 正确密码验证
		if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(p)); err != nil {
			c.JSON(401, gin.H{
				"state": "401",
				"message": "密码错误",
			})
		} else {
			t, err := auth.GenerateToken(a)
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
}

