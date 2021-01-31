package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hvxahv/pkg/utils"
	"log"
	"strings"
)

// JWTAuth ... 进行 JWT 验证的中间件
// 检测请求中是否携带 Token
// 判断 Token 是否正确
// 如果 Token 正确，将通过解析 Token 获取用户名并将用户名添加在上下文将 key 设置为 loginUser
func JWTAuth(c *gin.Context) {
	ht := c.Request.Header.Get("Authorization")
	t := strings.Split(ht, "Bearer ")[1]
	// 如果用户 header 请求中未携带 token
	if ht == "" {
		c.JSON(500, gin.H{
			"state": "500",
			"message": "请求中未携带 Token ",
		})
		c.Abort()
		return
	}
	_, _, err := JwtParseToken(t)
	if err != nil {
		c.JSON(500, gin.H{
			"state": "500",
			"message": "登陆失败 Token 不正确！",
		})
		c.Abort()
	} else {
		u, err := JwtParseUser(t)
		if err != nil {
			log.Println("通过 token 获取用户失败")
		}
		c.Set("loginUser", u.User)
		c.Next()
	}

}

func JwtParseToken(tokenString string) (*jwt.Token, *utils.Claims, error) {
	Claims := &utils.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return utils.K, nil
		})
	if err != nil {
		log.Println("解 Token 失败！")
	}
	return token, Claims, err
}

func JwtParseUser(tokenString string) (*utils.Claims, error) {
	if tokenString == "" {
		log.Println("需要传 Token ")
	}
	Claims := &utils.Claims{}
	_, err := jwt.ParseWithClaims(tokenString, Claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(utils.K), nil
		})
	if err != nil {
		return nil, err
	}
	return Claims, err
}
