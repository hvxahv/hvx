package notify

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/push"
)

func TestPush(t *testing.T) {
	d, err := json.Marshal(push.NewData("Notify", "Life's Not Out To Get You!", "https://avatars.githubusercontent.com/u/94792300?s=200&v=4", "Normal"))
	if err != nil {
		log.Println(err)
		return
	}
	if err := NewPush(724349838275870721, d).Push(); err != nil {
		fmt.Println(err)
		return
	}
}
