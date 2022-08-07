package internal

import "testing"

func TestAdministrates_IsChannelOwner(t *testing.T) {
	owner := NewAdministratesPermission(785746792214626305, 785518574097694721).IsChannelOwner()
	t.Log(owner)

	owner2 := NewAdministratesPermission(785746792214626305, 785747557487738881).IsChannelOwner()
	t.Log(owner2)
}

func TestNewAdministratesAdd(t *testing.T) {
	if err := NewAdministratesAdd(785746792214626305, 785747557487738881).AddAdministrator(); err != nil {
		t.Error(err)
		return
	}
}
