package errors

import "github.com/gin-gonic/gin"

const (
	ErrInternalServer        = "INTERNAL_SERVER_ERROR"
	ErrNotAuthorizationTOKEN = "AUTHORIZATION_TOKEN_IS_NOT_IN_REQUEST"
	ErrTokenParse            = "TOKEN_PARSING_FAILURE"
)

func NewHandler(code, errors string) gin.H {
	return gin.H{
		"code":   code,
		"errors": errors,
	}
}
