package jwt

import (
	"github.com/spf13/viper"
	"testing"
	"time"
)

func TestClaims_JWTTokenGenerator(t *testing.T) {
	e := time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour
	nc := NewClaims("x@disism.com", "123123", "hvturingga", "1-2-3-4-5", e)
	generator, err := nc.JWTTokenGenerator(viper.GetString("authentication.token.signed"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(generator)
}
