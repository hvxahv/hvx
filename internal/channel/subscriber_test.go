package channel

import (
	"fmt"
	"testing"
)

func TestNewSubscriber(t *testing.T) {
	TestInitChannelConfig(t)
	//
	//sub := "https://halfmemories.com/u/alice/inbox"
	//ns, err := NewSubscribers(692910076694757377, "alice", sub)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if err := ns.Create(); err != nil {
	//	fmt.Println(err)
	//}

}

func TestSubscribes_GetLisByID(t *testing.T) {
	TestInitChannelConfig(t)

	ns, err := NewGetSubscribersID(692668434193383425, 692283236803346433)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ns)

	r, err := ns.GetSubscribersByID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
