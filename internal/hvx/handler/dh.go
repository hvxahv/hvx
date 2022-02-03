package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
)

func DHRequestHandler(c *gin.Context) {
	client, err := account.NewECDHClient()
	if err != nil {
		return
	}

	d := &pb.NewDHRequestEncryption{
		Username:    middleware.GetUsername(c),
		DeviceId:    c.PostForm("device_id"),
		To:          c.PostForm("to"),
		DhPublicKey: c.PostForm("dh_public_key"),
		DhIv:        c.PostForm("dh_iv"),
	}
	reply, err := client.DHRequestEncryption(c, d)
	c.JSON(200, gin.H{
		"code":        reply.Code,
		"private_key": reply.Reply,
	})
}

func DHGetPublicHandler(c *gin.Context) {
	client, err := account.NewECDHClient()
	if err != nil {
		return
	}
	d := &pb.NewDHGetPublic{
		DeviceId: c.Param("id"),
	}
	public, err := client.DHGetPublic(c, d)
	if err != nil {
		return
	}
	fmt.Println(public)
	c.JSON(200, gin.H{
		"code":       public.Code,
		"device_id":  public.DeviceId,
		"iv":         public.Iv,
		"public_key": public.PublicKey,
	})
}

func DHSendHandler(c *gin.Context) {
	client, err := account.NewECDHClient()
	if err != nil {
		return
	}
	d := &pb.NewDHSendEncryption{
		DeviceId:    c.PostForm("device_id"),
		DhPublicKey: c.PostForm("dh_public_key"),
		PrivateKey:  c.PostForm("private_key"),
	}
	encryption, err := client.DHSendEncryption(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  encryption.Code,
		"reply": encryption.Reply,
	})
}

func DHWaitHandler(c *gin.Context) {
	client, err := account.NewECDHClient()
	if err != nil {
		return
	}
	d := &pb.NewDHWaitEncryption{DeviceId: c.Param("id")}
	encryption, err := client.DHWaitEncryption(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":          encryption.Code,
		"dh_public_key": encryption.DhPublicKey,
		"private_key":   encryption.PrivateKey,
	})
}
