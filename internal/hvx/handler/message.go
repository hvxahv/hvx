package handler

import (
	"github.com/gin-gonic/gin"
)

func GetMessageAccessHandler(c *gin.Context) {
	//name := middleware.GetUsername(c)

	//matrix, err := message.NewMatricesAccountID(name).Get()
	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"code":   "401",
	//		"matrix": "UNREGISTERED",
	//	})
	//	return
	//}

	//c.JSON(200, gin.H{
	//	"code":   "200",
	//	"matrix": matrix,
	//})
}

//
//func NewMessagesAccessHandler(c *gin.Context) {
//	name := middleware.GetUsername(c)
//	password := c.PostForm("password")
//
//	// The user is required to enter the password again to confirm the legitimacy of the account.
//	_, _, err := account.NewAuth(name, password).SignIn()
//	if err != nil {
//		log.Println(err)
//		if err.Error() == "PASSWORD_VERIFICATION_FAILED" {
//			c.JSON(401, gin.H{
//				"code":    "401",
//				"message": "PASSWORD_VERIFICATION_FAILED",
//			})
//			return
//		}
//		c.JSON(401, gin.H{
//			"code":    "401",
//			"message": "USERNAME_OR_PASSWORD_ERROR",
//		})
//		return
//	}
//
//	if err := message.NewMatrixAccessAuth(name, password).Register(); err != nil {
//		return
//	}
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"message": "ok!",
//	})
//}
