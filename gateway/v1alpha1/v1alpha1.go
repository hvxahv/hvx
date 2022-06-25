package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/pkg/proxy"
)

func SearchActorsHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/search/"+c.Param("actor"), "http://hvxahv.disism.internal:7042").Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}

func DeleteAccountHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/account", "http://hvxahv.disism.internal:7042").Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}
