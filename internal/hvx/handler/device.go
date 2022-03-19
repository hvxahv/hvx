/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/pkg/identity/middleware"
)

func GetDevices(c *gin.Context) {
	id := middleware.GetAccountID(c)
	client, err := device.GetDeviceClient()
	if err != nil {
		return
	}
	devices, err := client.GetDevicesByAccountID(c, &v1alpha1.GetDevicesByAccountIDRequest{
		AccountId: id,
	})
	if err != nil {
		return
	}
	c.JSON(200, devices)
}

func DeleteDevice(c *gin.Context) {
	id := middleware.GetAccountID(c)
	client, err := device.GetDeviceClient()
	if err != nil {
		return
	}
	reply, err := client.DeleteDeviceByID(c, &v1alpha1.DeleteDeviceByIDRequest{
		AccountId: id,
		DeviceId:  c.Param("id"),
	})
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":  "200",
		"reply": reply.Reply,
	})
}
