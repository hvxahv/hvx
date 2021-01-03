package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"hvxahv/app/accounts/app"
	"time"
)

var K = []byte("jwt_key_godis.disism.com")

type Claims struct {
	ID uint
	User string
	jwt.StandardClaims
}

// GenerateToken ...
func GenerateToken(u app.Accounts) (string, error) {
	expireTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		ID:  u.ID,
		User: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "token",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(K)
	if err != nil {
		fmt.Println(err)
	}

	return token, nil
}