package notify

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/push"
)

func TestPush(t *testing.T) {
	data := push.NewPushData("Notify", "Life's Not Out To Get You", "https://avatars.githubusercontent.com/u/94792300?s=200&v=4", "Normal")
	d, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	if err := NewPush(3812738129491231221, d).Push(); err != nil {
		fmt.Println(err)
		return
	}
}
