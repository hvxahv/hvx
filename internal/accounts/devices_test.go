package accounts

import (
	"fmt"
	"testing"
)

func TestDevices_Create(t *testing.T) {

}

func TestDevices_IsNotExist(t *testing.T) {

	b := NewDevicesID("74f435a8-df51-4816-9a9a-cb34688f68f4").IsNotExist()
	fmt.Println(b)
}

func TestDevices_DeleteByAccountID(t *testing.T) {

	if err := NewDevicesByAccountID(714483855961915393).DeleteALLByAccountID(); err != nil {
		fmt.Println(err)
		return
	}
}
