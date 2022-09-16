package internal

import (
	"github.com/hvxahv/hvx/cfg"
	"golang.org/x/net/context"
	"testing"
)

func init() {
	cfg.Default()
}

func TestAuthorization_Authorization(t *testing.T) {
	ctx := context.Background()

	v, err := NewAuthorization(ctx).Authorization("hvturingga", "hvxahv123")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
