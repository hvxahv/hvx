package notify

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/push"
)

func TestPush(t *testing.T) {
	d, err := json.Marshal(push.NewData("Notify",
		"Life's Not Out To Get You!",
		"https://avatars.githubusercontent.com/u/94792300?s=200&v=4",
		"Authorized"))
	if err != nil {
		log.Println(err)
		return
	}
	if err := NewPush(725220233314041857, 726287557594087425, d).Push(); err != nil {
		fmt.Println(err)
		return
	}
}
