package db

import (
	"testing"
)

func TestInitDB(t *testing.T) {

	nd := NewDb()
	if err := nd.InitDB(); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}

func TestCreateDB(t *testing.T) {
	nd := NewDb()

	name := "hvxahv"

	if err := nd.Create(name); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}
