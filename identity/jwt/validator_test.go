package jwt

import (
	"github.com/spf13/viper"
	"testing"
)

func TestValidator_Verify(t *testing.T) {
	verify, err := NewValidator(token, viper.GetString("authentication.token.signed")).Verify()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(verify)
}

func TestValidator_Authenticate(t *testing.T) {
	authenticate, err := NewValidator(token, viper.GetString("authentication.token.signed")).Authenticate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(authenticate)
}
