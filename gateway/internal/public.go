package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/gateway/address"
	"github.com/hvxahv/hvx/gateway/proxy"
)

func PublicHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/public"+c.Param("x"), address.GetHTTP(address.Public)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func WellKnownHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/.well-known/webfinger", address.GetHTTP(address.Public)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}

func GetActorHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/u/"+c.Param("actor"), address.GetHTTP(address.Public)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}

func AuthHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/auth", address.GetHTTP(address.Auth)).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}
