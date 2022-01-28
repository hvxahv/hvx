package device

import (
	"fmt"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cache"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"os"
	"testing"
)

func init() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	n := cockroach.NewDBAddr()
	if err := n.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
		return
	}

	cache.InitRedis(1)
}

func TestDevice_GetDevicesByAccountID(t *testing.T) {
	d := &v1alpha1.NewAccountID{AccountId: "731607090811043841"}
	s := device{}
	devices, err := s.GetDevicesByAccountID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(devices.Code, devices)
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
