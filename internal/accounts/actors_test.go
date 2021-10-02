package accounts

import (
	"fmt"
	"log"
	"testing"
)

func TestActors_NewActor(t *testing.T) {
	TestInitDB(t)
	
	a := NewActors("hvturingga", "", "")
	actor, err := a.NewActor()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(actor)
}

func TestActors_FindActorByPreferredUsername(t *testing.T) {
	TestInitDB(t)

	a := NewActorsPreferredUsername("hvturingga")

	r, err := a.FindByPreferredUsername()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r)
}

func TestActors_FindActorByID(t *testing.T) {
	TestInitDB(t)

	a := NewActorID(696901249207894017)
	actor, err := a.FindActorByID()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(actor)
}

func TestActors_Update(t *testing.T) {
	TestInitDB(t)

	a := NewActorID(698320974363721729)

	a.ActorType = ""
	a.Name = "HVTURINGGA"
	a.Summary = "我正在拍摄一个短片，欢迎关注我的 YouTube 频道...."
	a.Avatar = "http://localhost:9000/avatar/6e24d0b3-cc27-425f-84d0-e9bda6d41014-Kobayashi-YUI.jpg"

	err := a.Update()
	if err != nil {
		log.Println(err)
		return 
	}
}