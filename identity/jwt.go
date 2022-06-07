package identity

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Claims Generate the data needed for TOKEN.
type Claims struct {
	Uuid     string
	Email    string
	ID       string
	Username string
	DeviceID string
	jwt.StandardClaims
}

func NewClaims(id, mail, username, deviceID string) *Claims {
	expireTime := time.Now().Add(time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour)
	c := &Claims{
		Email:    mail,
		ID:       id,
		Username: username,
		DeviceID: deviceID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    viper.GetString("domain"),
			Subject:   "token",
		},
	}
	return c
}

// GenToken After the user logs in and the password is successfully verified,
// this method will be used to generate a Token and return.
func GenToken(id, mail, username, deviceID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(id, mail, username, deviceID))
	token, err := t.SignedString([]byte(viper.GetString("authentication.token.signed")))
	if err != nil {
		fmt.Println(err)
	}
	return token, nil
}

// VerifyToken Used to verify the validity of the TOKEN.
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("authentication.token.signed")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ParseToken Parse the data contained in TOKEN.
func ParseToken(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, errors.Errorf("DID NOT GET TOKEN!")
	}
	c := &Claims{}
	_, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), c,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("authentication.token.signed")), nil
		})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c, err
}