package activity

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestMessages_Outbox(t *testing.T) {
	IniTestConfig(t)

	// Prepare the data first.
	nf := NewFollowRequest("hvturingga", "https://mas.to/users/hvturingga")
	data, err := json.Marshal(nf)
	if err != nil {
		log.Println(err)
		return
	}

	nar := NewActivityRequest(nf.Actor, nf.Object, data, []byte(getPrivk()))
	nar.Follow()

}

func TestNewCreateNote(t *testing.T) {
	IniTestConfig(t)

	ncn := NewCreateNote()
	data, err := json.Marshal(ncn)
	if err != nil {
		log.Println(err)
		return
	}

	nar := NewActivityRequest(ncn.Actor, "https://mas.to/users/hvturingga", data, []byte(getPrivk()))
	nar.Create()
}

func TestNewAccept(t *testing.T) {
	IniTestConfig(t)

	name := "hvturingga"
	actor := "https://mas.to/users/hvturingga"
	oid := "https://mas.to/2db62f3e-3663-4be2-b881-1576bdf0e279"
	object := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)

	na := NewFollowAccept(name, actor, oid)

	data, err := json.Marshal(na)
	if err != nil {
		log.Println(err)
		return
	}

	nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
	nar.Accept()
}

// reply
/*
{
	"@context": "https://www.w3.org/ns/activitystreams",

	"id": "https://my-example.com/create-hello-world",
	"type": "Create",
	"actor": "https://my-example.com/actor",

	"object": {
		"id": "https://my-example.com/hello-world",
		"type": "Note",
		"published": "2018-06-23T17:17:11Z",
		"attributedTo": "https://my-example.com/actor",
		"inReplyTo": "https://mastodon.social/@Gargron/100254678717223630",
		"content": "<p>Hello world</p>",
		"to": "https://www.w3.org/ns/activitystreams#Public"
	}
}
 */