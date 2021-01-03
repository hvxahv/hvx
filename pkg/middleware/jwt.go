package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hvxahv/pkg/auth"
	"log"
	"strings"
)

// JWTAuth ...
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
		log.Println(u.User)
		c.Set("loginUser", u.User)
		c.Next()
	}

}

func JwtParseToken(tokenString string) (*jwt.Token, *auth.Claims, error) {
	Claims := &auth.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return auth.K, nil
		})
	if err != nil {
		log.Println("解 Token 失败！")
	}
	return token, Claims, err
}

func JwtParseUser(tokenString string) (*auth.Claims, error) {
	if tokenString == "" {
		log.Println("需要传 Token ")
	}
	Claims := &auth.Claims{}
	_, err := jwt.ParseWithClaims(tokenString, Claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(auth.K), nil
		})
	if err != nil {
		return nil, err
	}
	return Claims, err
}
