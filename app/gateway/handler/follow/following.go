package follow

import (
	"github.com/gin-gonic/gin"
	"hvxahv/pkg/activitypub"
	"hvxahv/pkg/mw"
)

// GetFollowingHandler ... 获取正在关注的用户
func GetFollowingHandler(c *gin.Context) {
	name := mw.GetUserName(c)
	r := activitypub.GetFollow(name, "following")
	c.JSON(200, gin.H{
		"res": r,
		"count": len(r),
	})
}
