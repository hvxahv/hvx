package device

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cache"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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

func TestDevice_Create(t *testing.T) {
	hash := uuid.New().String()
	d := &v1alpha1.NewDeviceCreate{
		AccountId: "731607090811043841",
		Ua:        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		Hash:      hash,
	}
	s := device{}
	create, err := s.Create(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(create.PublicKey, d.Hash)
}

func TestDevice_GetDevicesByAccountID(t *testing.T) {
	d := &v1alpha1.NewDeviceAccountID{AccountId: "731607090811043841"}
	s := device{}
	devices, err := s.GetDevicesByAccountID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(devices.Code, devices)
}

func TestDevice_GetDevicesByHash(t *testing.T) {
	d := &v1alpha1.NewDeviceHash{Hash: "d76cb739-f885-4793-b320-6b10c2d40f9b"}
	s := device{}
	v, err := s.GetDeviceByHash(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v.Hash, v.PublicKey, v.AccountId, v.Device, v.Id, v.PrivateKey)
}

func TestDevice_GetDeviceByID(t *testing.T) {
	d := &v1alpha1.NewDeviceID{Id: "732283914892804097"}
	s := device{}
	v, err := s.GetDeviceByID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v.Hash, v.PublicKey, v.AccountId, v.Device, v.Id, v.PrivateKey)
}

func TestDevice_DeleteAllByAccountID(t *testing.T) {
	d := &v1alpha1.NewDeviceAccountID{AccountId: "731354671656108033"}
	s := device{}
	reply, err := s.DeleteAllByAccountID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(reply.Code, reply.Reply)
}

func TestDevice_Delete(t *testing.T) {
	d := &v1alpha1.NewDeviceID{Id: "731877891266805761"}
	s := device{}
	reply, err := s.Delete(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(reply.Code, reply.Reply)
}

func TestDevice_IsExist(t *testing.T) {
	d := &v1alpha1.NewDeviceHash{Hash: "86d9d38d-ae81-4279-b36e-d22ecb26e35e"}
	s := device{}
	exist, err := s.IsExist(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(exist.IsExist)
}

func TestDevice_DeleteByDeviceHash(t *testing.T) {
	d := &v1alpha1.NewDeviceHash{Hash: "06e08310-c421-4ffe-98aa-f8f759e9f094"}
	s := device{}
	reply, err := s.DeleteByDeviceHash(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(reply.Code, reply.Reply)
}
