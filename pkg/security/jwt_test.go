package security

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken("foo", "bar", "123")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestVerifyToken(t *testing.T) {
	token, err := VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVdWlkIjoiZm9vIiwiVXNlciI6ImJhciIsImV4cCI6MTYzMDg0OTczMywiaWF0IjoxNjI4MjU3NzMzLCJpc3MiOiJsb2NhbGhvc3QiLCJzdWIiOiJ0b2tlbiJ9.gxIUL5JBzlAr1sHutaCNAVuNt_5oa9sqSs8n5mgmc88")
	if err != nil {
		t.Errorf("token parsing error: %v", err)
	}
	t.Log(token)
	
	
}

func TestParseToken(t *testing.T) {
	token, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVdWlkIjoiZm9vIiwiVXNlciI6ImJhciIsImV4cCI6MTYzMDg0OTczMywiaWF0IjoxNjI4MjU3NzMzLCJpc3MiOiJsb2NhbGhvc3QiLCJzdWIiOiJ0b2tlbiJ9.gxIUL5JBzlAr1sHutaCNAVuNt_5oa9sqSs8n5mgmc88")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(token.User)
}

func TestNewClaims(t *testing.T) {

	NewClaims("foo", "bar", "123")
}