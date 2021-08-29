package security

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	randomString, err := GenerateRandomString(15)
	if err != nil {
		t.Error(err)
	}
	t.Log(randomString)
}
