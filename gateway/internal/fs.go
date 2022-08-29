package internal

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/fs"
)

func AvatarHandler(c *gin.Context) {
	f, err := c.FormFile("avatar")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	}

	if f.Size > 1024*1024*2 {
		c.JSON(403, gin.H{
			"code":   "403",
			"errors": "FILE_SIZE_OUT_OF_RANGE",
		})
		return
	}

	ft := f.Header.Get("Content-Type")
	// Only image formats of type PNG / JPEG are accepted.
	if ft != "image/png" && ft != "image/jpeg" {
		c.JSON(403, gin.H{
			"code":   "403",
			"errors": "FORMAT_NOT_OF_PNG_OR_JPEG",
		})
		return
	}
	client, err := fs.NewMinio().Dial()
	if err != nil {
		c.JSON(500, gin.H{
			"code":   "500",
			"errors": "INTERNAL_SERVER_ERROR",
		})
		return
	}
	reader, _ := f.Open()
	info, err := client.SetPutOption("avatar", f.Filename, ft).FilePut(reader, f.Size)
	if err != nil {
		c.JSON(500, gin.H{
			"code":   "500",
			"errors": "FAILED_TO_UPLOAD_FILE",
		})
		return
	}
	fmt.Println(info)
	c.JSON(200, gin.H{
		"code":   "200",
		"avatar": info.Location,
	})
}
