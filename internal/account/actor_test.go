package account

import (
	"fmt"
	"log"
	"testing"
)

func TestActors_Create(t *testing.T) {
	a := NewActors("hvturingga", "", "Person")
	actor, err := a.Create()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(actor)
}

func TestActors_Update(t *testing.T) {
	a := NewActorsID(730822679860805633).SetActorName("HVTURINGGA").SetActorSummary("Aha ...")
	log.Println(a.ID)
	if err := a.Update(); err != nil {
		t.Error(err)
		return
	}
}

func TestNewAddActors(t *testing.T) {
	a := NewAddActors("xxs", "halfmemories.com", "", "", "", "https://halfmemories.com/u/xxs/inbox", "https://halfmemories.com/u/xxs", "", "Services")
	actor, err := a.Create()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)
}

func TestActors_GetActorByAccountUsername(t *testing.T) {
	account, err := NewActorsAccountUsername("hvturingga").GetActorByAccountUsername()
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(account)
}

func TestActors_GetActorByID(t *testing.T) {
	actor, err := NewActorsID(730799720355201025).GetActorByID()
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(actor)
}

func TestActors_GetActorByAddress(t *testing.T) {
	actor, err := NewActorsAddress("https://hvxahv.halfmemories.com/u/hvturingga").GetActorByAddress()
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(actor)
}
