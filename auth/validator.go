package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

type validator struct {
	token  string
	signer string
}

func NewValidator(token string, signer string) *validator {
	return &validator{token: token, signer: signer}
}

type v interface {
	Verify() (*jwt.Token, error)
	Authenticate() (*Claims, error)
}

func (v *validator) Verify() (*jwt.Token, error) {
	parse, err := jwt.Parse(v.token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(v.signer), nil
	})
	if err != nil {
		return nil, err
	}
	return parse, nil
}

func (v *validator) Authenticate() (*Claims, error) {
	c := &Claims{}
	_, err := jwt.ParseWithClaims(strings.TrimPrefix(v.token, "Bearer "), c,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(v.signer), nil
		})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c, err
}
