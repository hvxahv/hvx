package jwt

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v4"
)

// https://datatracker.ietf.org/doc/html/rfc7519

type Jwt interface {
	JWTTokenGenerator(signer string) (token string, err error)
}

var now = time.Now

type Userdata struct {
	Mail      string `json:"mail,omitempty"`
	AccountId string `json:"account_id,omitempty"`
	ActorId   string `json:"actor_id,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
	Username  string `json:"username,omitempty"`
}

// Claims Generate the data needed for TOKEN.
type Claims struct {
	Userdata             `json:"userdata"`
	jwt.RegisteredClaims `json:"jwt_._standard_claims"`
}

// NewClaims ...
// Expires Example: time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour
func NewClaims(mail, accountId, actorId, username, deviceId string, expires time.Duration) *Claims {
	return &Claims{
		Userdata: Userdata{
			Mail:      mail,
			AccountId: accountId,
			ActorId:   actorId,
			Username:  username,
			DeviceID:  deviceId,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    fmt.Sprintf("%s-hvx-issues", deviceId),
			Subject:   fmt.Sprintf("%s:token", username),
			ExpiresAt: jwt.NewNumericDate(now().Add(expires)),
			NotBefore: jwt.NewNumericDate(now().UTC()),
			IssuedAt:  jwt.NewNumericDate(now().UTC()),
			ID:        uuid.New().String(),
		},
	}
}
