package public

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/pkg/proxy"
	"github.com/hvxahv/hvx/pkg/v"
)

func Handler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/public"+c.Param("x"), v.GetRestServiceAddress("public")).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func WellKnownHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/.well-known/webfinger", v.GetRestServiceAddress("public")).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}

func GetActorHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/u/"+c.Param("actor"), v.GetRestServiceAddress("public")).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}
