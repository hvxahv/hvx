package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hvxahv/hvx/errors"
)

type Userdata struct {
	AccountId string `json:"account_id,omitempty"`
	ActorId   string `json:"actor_id,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Mail      string `json:"mail,omitempty"`
}

func NewUserdata(accountId string, actorId string, deviceID string, username string, mail string) *Userdata {
	return &Userdata{AccountId: accountId, ActorId: actorId, DeviceID: deviceID, Username: username, Mail: mail}
}

type auth interface {
	JWTTokenGenerator(secret string) (token string, err error)
}

//JWTTokenGenerator ...
func (c *Claims) JWTTokenGenerator(secret string) (token string, err error) {
	g, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return g, nil
}

type JWTParse struct {
	Token  string
	Secret string
}

type v interface {
	JWTTokenParse() (*Claims, error)
}

func NewParseJWTToken(token string, secret string) *JWTParse {
	return &JWTParse{Token: token, Secret: secret}
}

func (p *JWTParse) JWTTokenParse() (*Claims, error) {
	parse, err := jwt.ParseWithClaims(p.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parse.Claims.(*Claims); ok && parse.Valid {
		return claims, nil
	}
	return nil, errors.New(errors.ErrTokenInvalid)
}
