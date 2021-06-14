package encrypt

import "testing"

func TestGenToken(t *testing.T) {
	token, err := GenToken("foo", "bar")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

