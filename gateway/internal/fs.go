package internal

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/fs"
	"log"

	"github.com/gin-gonic/gin"
)

func AvatarHandler(c *gin.Context) {
	f, err := c.FormFile("avatar")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	}

	if f.Size > 1024*1024*2 {
		c.JSON(403, errors.NewHandler("403", errors.ErrFilesSize))
		return
	}

	ft := f.Header.Get("Content-Type")
	var suffix string
	switch ft {
	case "image/png":
		suffix = "png"
	case "image/jpeg":
		suffix = "jpeg"
	default:
		c.JSON(403, errors.NewHandler("403", errors.ErrFilesAllowedPNGAndJPEG))
		return
	}
	fn := fmt.Sprintf("%s.%s", uuid.New().String(), suffix)
	reader, _ := f.Open()
	put, err := fs.NewFs("minio", "avatar", fn, ft, reader, f.Size).Put()
	if err != nil {
		log.Println(err)
		c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut))
		return
	}
	c.JSON(200, gin.H{
		"code":   "200",
		"avatar": put,
	})
}

func FsHandler(c *gin.Context) {

}
