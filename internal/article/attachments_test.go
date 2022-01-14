package article

import (
	"fmt"
	"testing"
)

func TestAttachments_Create(t *testing.T) {
	a := NewAttachments("KOBAYASHI YUI", "image/jpeg", "")
	if err := a.Create(); err != nil {
		fmt.Println(err)
		return
	}
}
