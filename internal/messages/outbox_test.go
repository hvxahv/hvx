package messages

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
