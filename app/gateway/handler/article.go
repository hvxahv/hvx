package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/app/gateway/client/social"
	"hvxahv/pkg/auth"
)

// CreateArticleHandler ...
func CreateArticleHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	social.CreateArticleClient(author)
}
// UpdateArticleHandler ...
func UpdateArticleHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	social.UpdateArticleClient(author)
}
// DeleteArticleHandler ...
func DeleteArticleHandler(c *gin.Context) {
	social.DeleteArticleClient("4124141241")
}

