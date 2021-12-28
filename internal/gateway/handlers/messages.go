package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/internal/messages"
	"log"
)

func GetMessageAccessHandler(c *gin.Context) {
	name := middleware.GetUsername(c)

	matrix, err := messages.NewMatricesAccountID(name).Get()
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"code":   "200",
		"matrix": matrix,
	})
}

func NewMessagesAccessHandler(c *gin.Context) {
	name := middleware.GetUsername(c)
	password := c.PostForm("password")

	// The user is required to enter the password again to confirm the legitimacy of the account.
	_, _, err := accounts.NewAuth(name, password).SignIn()
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "USERNAME_OR_PASSWORD_ERROR",
		})
		return
	}

	if err := messages.NewMatrixAccessAuth(name, password).Register(); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}
