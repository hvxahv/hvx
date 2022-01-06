package device

import (
	"fmt"
	"log"
	"testing"
)

func TestDevices_Create(t *testing.T) {

}

func TestDevices_IsNotExist(t *testing.T) {

	b := NewDevicesIsNotExist("74f435a8-df51-4816-9a9a-cb34688f68f4").IsNotExist()
	fmt.Println(b)
}

func TestDevices_GetDevice(t *testing.T) {
	device, err := NewDevicesByID(725154842727251969, 725154968683708417).GetDevice()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(device)
}

func TestNewDeviceByHash(t *testing.T) {
	device, err := NewDeviceByHash(725154842727251969, "c17c3cfb-16bf-4efc-8771-95dfa03617c5").GetDeviceByHash()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(device)
}
func TestDevices_GetDevicesByAccountID(t *testing.T) {
	devices, err := NewDevicesByAccountID(725154842727251969).GetDevicesByAccountID()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(devices)
}

func TestDevices_Delete(t *testing.T) {
	if err := NewDeviceByHash(123, "").Delete(); err != nil {
		log.Println(err)
		return
	}
}

func TestDevices_DeleteByAccountID(t *testing.T) {

	if err := NewDevicesByAccountID(714483855961915393).DeleteAll(); err != nil {
		fmt.Println(err)
		return
	}
}
