package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/gateway/address"
	"github.com/hvxahv/hvx/gateway/proxy"
)

func SearchActorsHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/search/"+c.Param("actor"), address.GetHTTP(address.Actor)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}

func AccountHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/account"+c.Param("x"), address.GetHTTP(address.Account)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ActorHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/actor", address.GetHTTP(address.Actor)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}

func DeviceHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/device"+c.Param("x"), address.GetHTTP(address.Device)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ChannelHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/channel"+c.Param("x"), address.GetHTTP(address.Channel)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ArticleHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/article"+c.Param("x"), address.GetHTTP(address.Article)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func SavedHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/saved"+c.Param("x"), address.GetHTTP(address.Saved)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ActivityHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/activity"+c.Param("x"), address.GetHTTP(address.Activity)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func MessageAccessHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/message/access"+c.Param("x"), address.GetHTTP(address.Message)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func AuthHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/auth"+c.Param("x"), address.GetHTTP(address.Auth)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}
