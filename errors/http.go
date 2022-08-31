package errors

import "github.com/gin-gonic/gin"

const (
	ErrInternalServer = "INTERNAL_SERVER_ERROR"
)

const (
	ErrFilesPut               = "FAILED_TO_UPLOAD_FILE"
	ErrFilesSize              = "FILE_SIZE_OUT_OF_RANGE"
	ErrFilesAllowedPNGAndJPEG = "FORMAT_NOT_OF_PNG_OR_JPEG"
)

func NewHandler(code, errors string) gin.H {
	return gin.H{
		"code":   code,
		"errors": errors,
	}
}
