package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// JWTAuth JWT authentication middleware for gin network framework,
// Check whether the Token is carried in the request, and verify whether the Token is correct,
// Will obtain the username by parsing the Token and add the username in the context and set the key to loginUser.
//func JWTAuth(c *gin.Context) {
//	ht := c.Request.Header.Get("Authorization")
//	t := strings.Split(ht, "Bearer ")[1]
//	if ht == "" {
//		c.JSON(500, gin.H{
//			"state":   "500",
//			"message": "Token is not carried in the request.",
//		})
//		c.Abort()
//		return
//	}
//	_, _, err := JwtParseToken(t)
//	if err != nil {
//		c.JSON(500, gin.H{
//			"state":   "500",
//			"message": "Login failed. Token is incorrect!",
//		})
//		c.Abort()
//	} else {
//		u, err := JwtParseUser(t)
//		if err != nil {
//			log.Println("通过 token 获取用户失败")
//		}
//		c.Set("loginUser", u.User)
//		c.Next()
//	}
//
//}
//
//func JwtParseToken(tokenString string) (*jwt.Token, *httputils2.Claims, error) {
//	Claims := &httputils2.Claims{}
//	token, err := jwt.ParseWithClaims(tokenString, Claims,
//		func(token *jwt.Token) (i interface{}, err error) {
//			return httputils2.K, nil
//		})
//	if err != nil {
//		log.Println("解 Token 失败！")
//	}
//	return token, Claims, err
//}
//
//func JwtParseUser(tokenString string) (*httputils2.Claims, error) {
//	if tokenString == "" {
//		log.Println("需要传 Token ")
//	}
//	Claims := &httputils2.Claims{}
//	_, err := jwt.ParseWithClaims(tokenString, Claims,
//		func(token *jwt.Token) (interface{}, error) {
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//			}
//			return []byte(httputils2.K), nil
//		})
//	if err != nil {
//		return nil, err
//	}
//	return Claims, err
//}



var k = []byte(viper.GetString("token_signed"))

type claims struct {
	Uuid   string
	User string
	jwt.StandardClaims
}

// GenToken ...
func GenToken(Uuid, username string) (string, error) {
	expireTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &claims{
		Uuid: Uuid,
		User: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "token",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(k)
	if err != nil {
		fmt.Println(err)
	}

	return token, nil
}