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
	d := &pb.NewAccountID{AccountId: middleware.GetAccountID(c)}
	devices, err := client.GetDevicesByAccountID(c, d)
	if err != nil {
		return
	}
	c.JSON(200, devices)
}
