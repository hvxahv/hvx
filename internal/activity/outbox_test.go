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

func TestNewAccept(t *testing.T) {
	IniTestConfig(t)

	name := "hvturingga"
	actor := "https://mas.to/users/hvturingga"
	oid := "https://mas.to/da60792e-4b56-4ee1-bb41-8ae60898346b"
	object := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)

	na := NewFollowAccept(name, actor, oid)

	data, err := json.Marshal(na)
	if err != nil {
		log.Println(err)
		return
	}

	// {
	// "@context":"https://www.w3.org/ns/activitystreams",
	// "id":"https://mas.to/users/hvturingga#accepts/follows/116195",
	// "type":"Accept",
	// "actor":"https://mas.to/users/hvturingga",
	// "object":{
	// 		"id":"https://f817-2408-832f-20b0-2310-f02c-8d6d-bf13-51df.ngrok.io/f0a8e9f3-3445-4f73-911c-07b3096fa0f9",
	// 		"type":"Follow",
	// 		"actor":"https://f817-2408-832f-20b0-2310-f02c-8d6d-bf13-51df.ngrok.io/u/hvturingga",
	// 		"object":"https://mas.to/users/hvturingga"
	// }
	// }

	// {
	// "@context":"https://www.w3.org/ns/activitystreams",
	// "id":"https://f817-2408-832f-20b0-2310-f02c-8d6d-bf13-51df.ngrok.io/u/hvturingga#accepts/follows/e23c891d-86b9-4374-9aa9-0b45bf9c93c9",
	// "type":"Accept",
	// "actor":"https://f817-2408-832f-20b0-2310-f02c-8d6d-bf13-51df.ngrok.io/u/hvturingga",
	// "object":{
	// 		"id":"https://mas.to/7335676f-529d-40d9-a307-cf165c13a32e",
	// 		"type":"Follow",
	// 		"actor":"https://mas.to/users/hvturingga",
	// 		"object":"https://f817-2408-832f-20b0-2310-f02c-8d6d-bf13-51df.ngrok.io/u/hvturingga"
	// }
	// }

	nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
	nar.Accept()

}
