package internal

import "testing"

func TestAdministrates_IsChannelOwner(t *testing.T) {
	owner := NewAdministratesPermission(786085616327884801, 785518573776797697).IsChannelOwner()
	t.Log(owner)

	owner2 := NewAdministratesPermission(786085616327884801, 785747557487738881).IsChannelOwner()
	t.Log(owner2)
}

func TestNewAdministratesAdd(t *testing.T) {
	//if err := NewAdministratesAdd(785746792214626305, 785747557487738881).AddAdministrator(); err != nil {
	//	t.Error(err)
	//	return
	//}
}
