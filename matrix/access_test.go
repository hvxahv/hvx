package matrix

import (
	"github.com/hvxahv/hvx/cfg"
	"testing"
)

func init() {

	cfg.Default()
}

func TestMatrix_RegisterDummy(t *testing.T) {
	res, err := New("").RegisterDummy("ameotoko2", "hvxahv123")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
	t.Log(res.UserID.String())
	t.Log(res.AccessToken)
	t.Log(res.ExpiresInMS)
	t.Log(res.DeviceID)
}
