package v1alpha

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/pkg/proxy"
	"github.com/hvxahv/hvx/pkg/v"
)

func SearchActorsHandler(c *gin.Context) {
	if err := proxy.NewProxy(c, "/api/v1/search/"+c.Param("actor"), v.GetRestServiceAddress("account")).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}
