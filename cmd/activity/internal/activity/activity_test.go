package activity

import (
	"fmt"
	"github.com/hvxahv/hvx/cfg"
	"testing"
)

func init() {
	cfg.DefaultConfig()
}

const (
	follow = `{"@context":"https://www.w3.org/ns/activitystreams","id":"https://mstdn.social/04f2a2dc-a054-4d1a-b87d-447b0affecc3","type":"Follow","actor":"https://mstdn.social/users/hvturingga","object":"https://halfmemories.com/u/hvturingga"}`
	undo   = `{"@context":"https://www.w3.org/ns/activitystreams","id":"https://mstdn.social/users/hvturingga#follows/751608/undo","type":"Undo","actor":"https://mstdn.social/users/hvturingga","object":{"id":"https://mstdn.social/04f2a2dc-a054-4d1a-b87d-447b0affecc3","type":"Follow","actor":"https://mstdn.social/users/hvturingga","object":"https://halfmemories.com/u/hvturingga"}}`
)

func TestInboxActivity_Activity(t *testing.T) {
	fmt.Println(follow)
	fmt.Println(undo)
}
