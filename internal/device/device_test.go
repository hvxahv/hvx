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
	"testing"
)

func TestAccount_DeviceIsExistByHash(t *testing.T) {
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

func TestAccount_CreateDevice(t *testing.T) {
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

func TestAccount_GetDevicesByAccountID(t *testing.T) {
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

func TestAccount_GetDeviceByID(t *testing.T) {
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

func TestAccount_GetDeviceByHash(t *testing.T) {
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

func TestAccount_DeleteAllByAccountID(t *testing.T) {
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

func TestAccount_DeleteDeviceByID(t *testing.T) {
	d := &v1alpha1.DeleteDeviceByIDRequest{
		DeviceId: "737991829363687425",
	}
	s := device{}
	del, err := s.DeleteDeviceByID(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(del)
}

func TestAccount_DeleteDeviceByHash(t *testing.T) {
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
