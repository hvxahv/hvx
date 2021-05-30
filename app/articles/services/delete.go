package services

//
//import (
//	"errors"
//	"github.com/mw-gonic/mw"
//	"log"
//)
//
//func DeleteArticleHandler(c *mw.Context) {
//	id := c.PostForm("id")
//	if err := DeleteArticleByID(id); err != nil {
//		log.Println("Delete Error!", err)
//	}
//	c.JSON(200, mw.H{
//		"status": "200",
//		"message": "DELETE SUCCESS",
//	})
//
//}
//
//func DeleteArticleByID(id string) error {
//	if err := db2.Debug().Table("status").Where("id = ?", id).Delete(&Article{}).Error; err != nil {
//		return errors.New("Delete Article Error")
//	}
//
//	return nil
//}
