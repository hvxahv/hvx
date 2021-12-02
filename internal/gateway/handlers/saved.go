package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func SavedUploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	open, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err := DetectContentType(open)
	if err != nil {
		return
	}
	fmt.Println(t)
}

// DetectContentType 判断文件类型返回字符串格式的文件类型名称。
func DetectContentType(out multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err2 := out.Read(buffer)
	if err2 != nil {
		return "", err2
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
