package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/disism/hvxahv/pkg/security"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// JWTAuth JWT authentication middleware for gin network framework,
// Check whether the Token is carried in the request, and verify whether the Token is correct,
// Will obtain the username by parsing the Token and add the username in the context and set the key to loginUser.
func JWTAuth(c *gin.Context) {
	ht := c.Request.Header.Get("Authorization")
	t := strings.Split(ht, "Bearer ")[1]
	if ht == "" {
		c.JSON(500, gin.H{
			"state":   "500",
			"message": "Token is not carried in the request.",
		})
		c.Abort()
		return
	}
	_, _, err := jwtParseToken(t)
	if err != nil {
		c.JSON(500, gin.H{
			"state":   "500",
			"message": "Login failed. Token is incorrect!",
		})
		c.Abort()
	} else {
		u, err := jwtParseUser(t)
		if err != nil {
			log.Println("Failed to obtain user through token.")
		}
		c.Set("loginUser", u.User)
		c.Next()
	}

}

func jwtParseToken(tokenString string) (*jwt.Token, *security.Claims, error) {
	Claims := &security.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return viper.GetString("token_signed"), nil
		})
	if err != nil {
		log.Println("Load verification token middleware.")
	}
	return token, Claims, err
}

func jwtParseUser(tokenString string) (*security.Claims, error) {
	if tokenString == "" {
		log.Println("Need to pass Token.")
	}
	Claims := &security.Claims{}
	_, err := jwt.ParseWithClaims(tokenString, Claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("token_signed")), nil
		})
	if err != nil {
		return nil, err
	}
	return Claims, err
}



