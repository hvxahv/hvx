package auth

import (
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v4"
)

// https://datatracker.ietf.org/doc/html/rfc7519

var now = time.Now

// Claims https://datatracker.ietf.org/doc/html/rfc7519#section-4
// JWT Claims.
type Claims struct {
	Userdata             `json:"userdata"`
	jwt.RegisteredClaims `json:"jwt_._standard_claims"`
}

func NewClaims(userdata *Userdata, registeredClaims *jwt.RegisteredClaims) *Claims {
	return &Claims{Userdata: *userdata, RegisteredClaims: *registeredClaims}
}

func NewRegisteredClaims(issuer, aud, sub string, expir time.Duration) *jwt.RegisteredClaims {
	return &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   sub,
		Audience:  []string{aud},
		ExpiresAt: jwt.NewNumericDate(now().Add(expir)),
		NotBefore: jwt.NewNumericDate(now().UTC()),
		IssuedAt:  jwt.NewNumericDate(now().UTC()),
		ID:        uuid.New().String(),
	}
}
