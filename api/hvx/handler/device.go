package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
)

func GetDevicesHandler(c *gin.Context) {
	client, err := device.NewDeviceClient()
	if err != nil {
		return
	}
	d := &pb.NewDeviceAccountID{AccountId: middleware.GetAccountID(c)}
	devices, err := client.GetDevicesByAccountID(c, d)
	if err != nil {
		return
	}
	c.JSON(200, devices)
}

func DeleteDevicesHandler(c *gin.Context) {
	hash := c.PostForm("device_hash")
	client, err := device.NewDeviceClient()
	if err != nil {
		return
	}
	d := &pb.NewDeviceHash{Hash: hash}
	reply, err := client.DeleteByDeviceHash(c, d)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":  reply.Code,
		"reply": reply.Reply,
	})
}
