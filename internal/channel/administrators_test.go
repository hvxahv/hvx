package channel

import (
	"fmt"
	"testing"
)

func TestNewAddAdmins(t *testing.T) {
	TestInitChannelConfig(t)

	na, err := NewAddAdmins(692910076694757377, 692283236803346433,  692635608323948545)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := na.Add(); err != nil {
		fmt.Println(err)
	}
}

func TestNewRemoveAdmByName(t *testing.T) {
	TestInitChannelConfig(t)

	nra, err := NewAddAdmins(692547656905850881, 692283236803346433, 692635286453518337)
	if err != nil {
		fmt.Println(err)
	}
	if err := nra.Remove(); err != nil {
		fmt.Println(err)
	}
}

func TestAdministrators_QueryAdmLisByID(t *testing.T) {
	TestInitChannelConfig(t)

	na := NewAdminsByID(692668434193383425, 692283236803346433)
	id, err := na.QueryAdmLisByCID()
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(id)
}