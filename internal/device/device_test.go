/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package device

import (
	"context"
	"fmt"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
}

func TestDevice_DeviceIsExistByHash(t *testing.T) {
	d := &v1alpha1.DeviceIsExistByHashRequest{
		Hash: "7126762b-3b46-441e-8428-54b5effe6bb9s",
	}
	s := device{}
	exist, err := s.DeviceIsExistByHash(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(exist)
}

func TestDevice_DeviceIsExistByID(t *testing.T) {
	d := &v1alpha1.DeviceIsExistByIDRequest{
		Id: "745875183290318849",
	}
	s := device{}
	exist, err := s.DeviceIsExistByID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(exist)
}

func TestDevice_CreateDevice(t *testing.T) {
	d := &v1alpha1.CreateDeviceRequest{
		AccountId: "733124680636596225",
		Ua:        "Edge",
		Hash:      "xx-xxx-x-xxx-xxx",
	}
	s := device{}
	c, err := s.CreateDevice(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(c)
}

func TestDevice_GetDevicesByAccountID(t *testing.T) {
	d := &v1alpha1.GetDevicesByAccountIDRequest{
		AccountId: "733124680636596225",
	}
	s := device{}
	devices, err := s.GetDevicesByAccountID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(devices)
}

func TestDevice_GetDeviceByID(t *testing.T) {
	d := &v1alpha1.GetDeviceByIDRequest{
		DeviceId: "737990596587618305",
	}
	s := device{}
	device, err := s.GetDeviceByID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(device)
}

func TestDevice_GetDeviceByHash(t *testing.T) {
	d := &v1alpha1.GetDeviceByHashRequest{
		Hash: "xx-xxx-x-xxx",
	}
	s := device{}
	device, err := s.GetDeviceByHash(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(device)
}

func TestDevice_DeleteAllByAccountID(t *testing.T) {
	d := &v1alpha1.DeleteDeviceAllByAccountIDRequest{
		AccountId: "733124680636596225",
	}
	s := device{}
	del, err := s.DeleteDeviceAllByAccountID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(del)
}

func TestDevice_DeleteDeviceByID(t *testing.T) {
	d := &v1alpha1.DeleteDeviceByIDRequest{
		AccountId: "",
		DeviceId:  "737991829363687425",
	}
	s := device{}
	del, err := s.DeleteDeviceByID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(del)
}

func TestDevice_DeleteDeviceByHash(t *testing.T) {
	d := &v1alpha1.DeleteDeviceByHashRequest{
		Hash: "xx-xxx-x-xxx-xxx",
	}
	s := device{}
	del, err := s.DeleteDeviceByHash(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(del)
}
