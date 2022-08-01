package auth

import "github.com/golang-jwt/jwt/v4"

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
	JWTTokenGenerator(signer string) (token string, err error)
}

//JWTTokenGenerator ...
func (c *Claims) JWTTokenGenerator(signer string) (token string, err error) {
	g, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(signer))
	if err != nil {
		return "", err
	}
	return g, nil
}
