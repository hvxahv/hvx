package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/saved"
	"log"
	"strconv"
)

func SavedHandler(c *gin.Context) {
	//a := middleware.GetUsername(c)
	//name := c.PostForm("name")
	//hash := c.PostForm("hash")
	//fileType := c.PostForm("type")
	//
	//account, err := account.NewAccountsUsername(a).GetAccountByUsername()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//if err := saved.NewSaves(account.ID, name, hash, fileType).Create(); err != nil {
	//	log.Println(err)
	//	return
	//}
	//c.JSON(200, gin.H{
	//	"code":    "200",
	//	"message": "ok",
	//})
}

func GetSavedByIDHandler(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}
	s, err := saved.NewSavesID(uint(i)).GetSavedByID()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": s,
	})
}

func GetSavesHandler(c *gin.Context) {
	//account, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//saves, err := saved.NewSavesByAccountID(account.ID).GetSaves()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//c.JSON(200, gin.H{
	//	"code":  "200",
	//	"saves": saves,
	//})
}
