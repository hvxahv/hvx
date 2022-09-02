package errors

import "github.com/gin-gonic/gin"

const (
	ErrInternalServer = "INTERNAL_SERVER_ERROR"
)

func NewHandler(code, errors string) gin.H {
	return gin.H{
		"code":   code,
		"errors": errors,
	}
}
