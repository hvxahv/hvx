package channel

import (
	"fmt"
	"testing"
)

func TestNewSubscriber(t *testing.T) {
	TestInitChannelConfig(t)

	ns, err := NewSubscribes(692668434193383425, "alice")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := ns.New(); err != nil {
		fmt.Println(err)
	}

}

func TestSubscribes_GetLisByID(t *testing.T) {
	TestInitChannelConfig(t)

	ns, err := NewSubLisByID(692668434193383425, 692283236803346433)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ns)

	r, err := ns.QueryLisByID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}