package internal

import (
	"testing"

	"github.com/hvxahv/hvx/cfg"
)

func init() {
	cfg.Default()
}

func TestActors_IsExist(t *testing.T) {

}

func TestActors_Create(t *testing.T) {

}

func TestActors_Get(t *testing.T) {

}

func TestActors_GetActorsByPreferredUsername(t *testing.T) {

}

func TestActors_Add(t *testing.T) {

}

func TestGetActorByUsername(t *testing.T) {

}

func TestActors_Edit(t *testing.T) {
	a := NewActorsId(1)
	if err := a.Edit(); err != nil {
		t.Error(err)
	}
}

func TestActors_Delete(t *testing.T) {
	a := NewActorsId(1)
	if err := a.Delete(); err != nil {
		t.Error(err)
	}
}
