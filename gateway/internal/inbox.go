package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/gateway/proxy"
)

func InboxHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/u/"+c.Param("actor")+"/inbox", public).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}
