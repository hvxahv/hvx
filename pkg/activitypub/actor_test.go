package activitypub

import (
	"fmt"
	"testing"
)

func TestGetActorName(t *testing.T) {
	t.Log(GetActorName("acct:hvturingga@fac47988a686.ngrok.io"))

	t.Log(GetActorName("acct:hvturingga"))
}

func TestGetHost(t *testing.T) {
	r := GetHost("acct:hvturingga@fac47988a686.ngrok.io")
	t.Log(r)

	noHost := GetHost("acct:hvturingga")
	t.Log(noHost)
}

func TestIsRemote(t *testing.T) {
	ir := IsRemote("acct:hvturingga@xxs.ngrok.io")
	fmt.Println(ir)
	ir2 := IsRemote("acct:hvturingga@fc3b47257006.ngrok.io")
	fmt.Println(ir2)
	ir3 := IsRemote("acct:hvturingga")
	fmt.Println(ir3)

}

func TestGetRemoteActor(t *testing.T) {
	actor, err := GetRemoteActor("https://mas.to/users/hvturingga")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(actor)
}