package httputils

import (
	"github.com/dgrijalva/jwt-go"
)

var K = []byte("jwt_key_hvxahv.half_memories.com")

type Claims struct {
	ID   uint
	User string
	jwt.StandardClaims
}

//// GenerateToken ...
//func GenerateToken(u accounts.Accounts) (string, error) {
//	expireTime := time.Now().Add(30 * 24 * time.Hour)
//	claims := &Claims{
//		ID:   u.ID,
//		User: u.Username,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expireTime.Unix(),
//			IssuedAt:  time.Now().Unix(),
//			Issuer:    "127.0.0.1",
//			Subject:   "token",
//		},
//	}
//	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := t.SignedString(K)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	return token, nil
//}
