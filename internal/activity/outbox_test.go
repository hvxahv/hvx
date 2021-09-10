package activity

import (
	"encoding/json"
	"log"
	"testing"
)

func TestMessages_Outbox(t *testing.T) {
	IniTestConfig(t)

	// Prepare the data first.
	nf := NewFollow("hvturingga", "https://mas.to/users/hvturingga")
	data, err := json.Marshal(nf)
	if err != nil {
		log.Println(err)
		return
	}

	nar := NewActivityRequest(nf.Actor, nf.Object, data, []byte(getPrivk()))
	nar.Follow()

}

func TestNewAccept(t *testing.T) {
	//	{
	//	"@context":"https://www.w3.org/ns/activitystreams",
	//	"id":"https://mas.to/09502b93-454f-4901-98af-4fa3a39d6427",
	// 	"type":"Follow",
	//	"actor":"https://mas.to/users/hvturingga",
	//	"object":"https://ef41-2408-832f-20b0-2310-d9a-998f-806b-4baf.ngrok.io/u/hvturingga"
	//	}
	name := "hvturingga"
	actor := "https://mas.to/users/hvturingga"
	oid := "https://mas.to/09502b93-454f-4901-98af-4fa3a39d6427"
	object := "https://ef41-2408-832f-20b0-2310-d9a-998f-806b-4baf.ngrok.io/u/hvturingga"

	IniTestConfig(t)
	na := NewAccept(name, actor, oid, object)

	data, err := json.Marshal(na)
	if err != nil {
		log.Println(err)
		return
	}

	// {
	// "@context":"https://www.w3.org/ns/activitystreams",
	// "id":"https://mas.to/users/hvturingga#accepts/follows/115804",
	// "type":"Accept",
	// "actor":"https://mas.to/users/hvturingga",
	// "object":{
	//		"id":"https://ef41-2408-832f-20b0-2310-d9a-998f-806b-4baf.ngrok.io/ccd2e75f-78aa-461d-b759-6219156f2327",
	//		"type":"Follow",
	//		"actor":"https://ef41-2408-832f-20b0-2310-d9a-998f-806b-4baf.ngrok.io/u/hvturingga",
	//		"object":"https://mas.to/users/hvturingga"
	//		}
	//	}

	nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
	nar.Accept()

}
