package accounts

import (
	"fmt"
	"log"
	"testing"
)

func TestActors_NewActor(t *testing.T) {
	a := NewActors("hvturingga", "", "", "Person")
	actor, err := a.NewActor()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(actor)
}

func TestActors_FindActorByPreferredUsername(t *testing.T) {
	a := NewActorsPreferredUsername("hvturingga")

	r, err := a.FindByPreferredUsername()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r)
}

func TestActors_FindActorByID(t *testing.T) {
	a := NewActorID(696901249207894017)
	actor, err := a.GetActorByUri()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(actor)
}

func TestActors_FindActorByUrl(t *testing.T) {
	a := NewActorUri("https://mas.to/users/hvturinggas")
	actor, err := a.GetActorByUri()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(actor)
}

func TestActors_Update(t *testing.T) {
	a := NewActorID(698619813575491585)

	a.ActorType = ""
	a.Name = "HVTURINGGA"
	a.Summary = "我正在拍摄一个短片，欢迎关注我的 YouTube 频道...."
	a.Avatar = "http://stage48.net/wiki/images/5/5b/KobayashiYui8th.jpg"

	err := a.Update()
	if err != nil {
		log.Println(err)
		return
	}
}
