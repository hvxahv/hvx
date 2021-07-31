package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// Claims ***************************************************
type Claims struct {
	Uuid   string
	User string
	jwt.StandardClaims
}

func NewClaims(uuid string, user string) *Claims {
	expireTime := time.Now().Add(30 * 24 * time.Hour)
	c := &Claims{
		Uuid: uuid,
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    viper.GetString("localhost"),
			Subject:   "token",
		},
	}
	return c
}

// GenToken After the user logs in and the password is successfully verified,
// this method will be used to generate a Token and return.
func GenToken(uuid, username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(uuid, username))
	token, err := t.SignedString([]byte(viper.GetString("token_signed")))
	if err != nil {
		fmt.Println(err)
	}

	return token, nil
}





