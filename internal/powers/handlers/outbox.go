package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetActorOutbox(c *gin.Context)  {
	name := c.Param("actor")

	fmt.Printf("Requested data for %s", name)

}
