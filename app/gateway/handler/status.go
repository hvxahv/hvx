package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/app/gateway/client/social"
	"hvxahv/pkg/auth"
)

/*
	Status Group
*/
func CreateStatusHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	social.CreateStatusClient(author)
}
// UpdateArticleHandler ...
func UpdateStatusHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	social.UpdateStatusClient(author)
}
// DeleteArticleHandler ...
func DeleteStatusHandler(c *gin.Context) {
	social.DeleteStatusClient("2321214241241")
}


