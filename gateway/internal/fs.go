package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/fs"
	"log"
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

	var (
		ft = f.Header.Get("Content-Type")
		fn = fmt.Sprintf("%s-%s", uuid.New().String(), f.Filename)
	)
	if ft != "image/png" && ft != "image/jpeg" {
		c.JSON(403, errors.NewHandler("403", errors.ErrFilesAllowedPNGAndJPEG))
		return
	}
	reader, _ := f.Open()
	put, err := fs.NewFsPut("minio", "avatar", fn, ft, reader, f.Size).Put()
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut+err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"code":   "200",
		"avatar": put,
	})
}

func AttachHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut))
		return
	}
	var idx int
	for i, _ := range form.File["attach"] {
		idx = i + 1
	}
	if idx > 9 {
		c.JSON(403, errors.NewHandler("403", errors.ErrMaximum))
		return
	}
	var attas []string
	for _, f := range form.File["attach"] {
		var (
			ft = f.Header.Get("Content-Type")
			fn = fmt.Sprintf("%s-%s", uuid.New().String(), f.Filename)
		)
		reader, _ := f.Open()
		put, err := fs.NewFsPut("minio", "attach", fn, ft, reader, f.Size).Put()
		if err != nil {
			c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut))
			return
		}

		attas = append(attas, put)
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"attas": attas,
	})
}
