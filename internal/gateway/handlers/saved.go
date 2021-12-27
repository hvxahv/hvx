package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/internal/saved"
	"log"
	"mime/multipart"
	"net/http"
)

func SavedHandler(c *gin.Context) {
	a := middleware.GetUsername(c)
	hash := c.PostForm("hash")
	fileType := c.PostForm("type")

	account, err := accounts.NewAccountsUsername(a).GetAccountByUsername()
	if err != nil {
		log.Println(err)
		return
	}
	if err := saved.NewSaves(account.ID, hash, fileType).Create(); err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok",
	})
}

// DetectContentType Judge the file type and return the file type name in string format.
func DetectContentType(out multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err2 := out.Read(buffer)
	if err2 != nil {
		return "", err2
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
