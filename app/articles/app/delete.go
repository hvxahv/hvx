package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func DeleteArticleHandler(c *gin.Context) {
	id := c.PostForm("id")
	if err := DeleteArticleByID(id); err != nil {
		log.Println("Delete Error!", err)
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "DELETE SUCCESS",
	})

}

func DeleteArticleByID(id string) error {
	if err := db2.Debug().Table("status").Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return errors.New("Delete Article Error")
	}

	return nil
}