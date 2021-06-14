package encrypt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var k = []byte("jwt_key_hvxahv.half_memories.com")

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

