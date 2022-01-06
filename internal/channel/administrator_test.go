package channel

import (
	"fmt"
	"testing"
)

func TestNewAddAdmins(t *testing.T) {
	TestInitChannelConfig(t)

	na, err := NewAddAdmins(701990339882680321, 698619813575491585, 699445624156061697)
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

	nra, err := NewAddAdmins(701990339882680321, 698619813575491585, 699445624156061697)
	if err != nil {
		fmt.Println(err)
	}
	if err := nra.Remove(); err != nil {
		fmt.Println(err)
	}
}

func TestAdministrators_QueryAdmLisByID(t *testing.T) {
	TestInitChannelConfig(t)

	na := NewISAdmins(701990339882680321, 698619813575491585)
	id, err := na.FindAdmLisByChannelID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}
