package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"log"
)

func GetPublicAccountCountHandler(c *gin.Context) {
	db := cockroach.GetDB()
	var count int64
	if err := db.Debug().Table("accounts").Count(&count).Error; err != nil {
		log.Println(err)
		return
	}
	fmt.Println(count)
	c.JSON(200, gin.H{
		"code":          "200",
		"account_count": count,
	})
}
